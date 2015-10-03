package card

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func openDB() (string, error) {
	dir, err := ioutil.TempDir("", "fowtcg-test")
	if err != nil {
		panic(fmt.Errorf("create temp dir failed: %s", err))
	}
	if err := OpenDB(dir); err != nil {
		return dir, fmt.Errorf("open db failed: %s", err)
	}
	return dir, nil
}

func TestPs2pi(t *testing.T) {
	p := []string{"a", "b", "c", "d", "e"}
	q := ps2pi(p)
	if len(q) != len(p) {
		t.Fatalf("have len %v, want %v", len(q), len(p))
	}
	for i, e := range q {
		if e.(string) != p[i] {
			t.Fatalf("have %s at %v, want %s", e.(string), i, p[i])
		}
	}
}

func TestDatabase(t *testing.T) {
	dir, err := openDB()
	defer os.RemoveAll(dir)
	if err != nil {
		t.Fatal(err)
	}

	var count int
	db.QueryRow("select count(*) from card;").Scan(&count)
	if count == 0 {
		t.Fatal("card table is empty")
	}

	if err := db.QueryRowx("select * from card where _id=1;").StructScan(&Card{}); err != nil {
		t.Fatalf("struct scan failed: %s", err)
	}

	cards, err := FindByTitle("alice")
	if err != nil {
		t.Fatal(err)
	}
	if len(cards) == 0 {
		t.Fatal("result length of zero")
	}
	for _, c := range cards {
		t.Logf("%#v\n", c)
	}
}
