package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	cg "github.com/alnah/go-auth/config"
	db "github.com/alnah/go-auth/db/dsn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := cg.LoadConfig("../../.env/", "test")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	dsn, err := db.GenerateDSN(db.GenerateDSNParams{
		User:     config.PostgresUser,
		Password: config.PostgresPassword,
		Host:     config.PostgresHost,
		Port:     config.PostgresPort,
		DBName:   config.PostgresName,
	})
	if err != nil {
		log.Fatal("cannot generate DSN:", err)
	}

	testDB, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("cannot connect to the database:", err)
	}

	defer func() {
		if err := testDB.Close(); err != nil {
			log.Fatal("cannot close the database connection:", err)
		}
	}()

	testQueries = New(testDB)

	test := m.Run()
	os.Exit(test)
}
