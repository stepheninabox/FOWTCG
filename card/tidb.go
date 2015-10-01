package card

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/jmoiron/sqlx"
	"github.com/ngaut/log"
	_ "github.com/pingcap/tidb"
)

var db *sqlx.DB

// OpenDB opens the database in the given directory; if required, the database
// will be initialized.
func OpenDB(dir string) (err error) {
	log.SetLevelByString("error")

	dir = filepath.Join(dir, "tidb")
	_, err = os.Stat(dir)
	init := os.IsNotExist(err)

	name := "tidb"
	db, err = sqlx.Open("tidb", "goleveldb://"+dir+"/"+name)
	if err != nil {
		return fmt.Errorf("db open failed: %s", err)
	}
	if err := db.Ping(); err != nil {
		return fmt.Errorf("db ping failed: %s", err)
	}

	if init {
		b, err := ioutil.ReadFile(filepath.Join("data", "schema.sql"))
		if err != nil {
			return fmt.Errorf("read schema file failed: %s", err)
		}
		if _, err = db.Exec(string(b)); err != nil {
			return fmt.Errorf("exec schema failed: %s", err)
		}
		if err = insertCards(); err != nil {
			return fmt.Errorf("insert cards failed: %s", err)
		}
	}

	return nil
}

// Transaction manages a db transaction, automatically calling commit
// on success or rollback on failure.
func Transaction(fn func(*sqlx.Tx) error) (err error) {
	var tx *sqlx.Tx
	tx, err = db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	return fn(tx)
}

// ps2pi converts []string to []interface
func ps2pi(p []string) []interface{} {
	q := make([]interface{}, len(p))
	for i, e := range p {
		q[i] = e
	}
	return q
}

// insertCards reads data from csv and inserts into database.
func insertCards() error {
	f, err := os.Open(filepath.Join("data", "cards.csv"))
	if err != nil {
		return fmt.Errorf("open cards.csv failed: %s", err)
	}
	defer f.Close()

	return Transaction(func(tx *sqlx.Tx) error {
		stmt, err := tx.Prepare("insert into card values(NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);")
		if err != nil {
			return fmt.Errorf("prepare statement failed: %s", err)
		}

		r := csv.NewReader(f)
		r.Comma = '\t'
		r.FieldsPerRecord = 12
		r.LazyQuotes = true
		// discard first line of field names
		_, _ = r.Read()
		idx := 0
		for rec, err := r.Read(); err != io.EOF; rec, err = r.Read() {
			if err != nil {
				return fmt.Errorf("read row(%v) failed: %s", idx, len(rec))
			}
			if _, err := stmt.Exec(ps2pi(rec)...); err != nil {
				return fmt.Errorf("insert record(%v) failed: %s: row: %#v", idx, err, rec)
			}
			idx++
		}

		return nil
	})
}
