package database_test

import (
	"log"
	"os"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jmoiron/sqlx"
)

type TestInfra struct {
	db *sqlx.DB
}

func TestMain(m *testing.M) {
	run := m.Run()

	os.Exit(run)
}

func NewTestInfra() *TestInfra {
	dbURI := "postgres://test_user:test_password@localhost:5432/test_db?sslmode=disable"

	db, err := sqlx.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Cannot open order DB: %v", err)
	}

	return &TestInfra{db: db}
}

func (t *TestInfra) GetDB() *sqlx.DB {
	return t.db
}

func hardDeleteForLocalTesting(db *sqlx.DB) error {
	db.Exec("ALTER TABLE transactions DROP CONSTRAINT transactions_source_account_id_fkey")
	db.Exec("ALTER TABLE transactions DROP CONSTRAINT transactions_destination_account_id_fkey")

	_, err := db.Exec("TRUNCATE table accounts")

	if err != nil {
		return err
	}

	_, err = db.Exec("TRUNCATE table transactions")

	if err != nil {
		return err
	}

	return nil
}
