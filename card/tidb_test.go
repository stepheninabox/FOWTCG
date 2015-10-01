package card

import (
	"io/ioutil"
	"os"
	"testing"
)

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
	dir, err := ioutil.TempDir("", "fowtcg-test")
	if err != nil {
		t.Fatalf("create temp dir failed: %s", err)
	}
	defer os.RemoveAll(dir)

	if err := OpenDB(dir); err != nil {
		t.Fatalf("open db failed: %s", err)
	}

	var count int
	db.QueryRow("select count(*) from card;").Scan(&count)
	if count == 0 {
		t.Fatal("card table is empty")
	}

	if err := db.QueryRowx("select * from card where _id=1;").StructScan(&Card{}); err != nil {
		t.Fatalf("struct scan failed: %s", err)
	}
}
