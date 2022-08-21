package models

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func newTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "postgres://test_web:pass@localhost:5432/test_tsnippet")
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("./testdata/setup.sql")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(string(script))
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		script, err := os.ReadFile("./testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(string(script))
		if err != nil {
			t.Fatal(err)
		}

		db.Close()
	})

	return db
}
