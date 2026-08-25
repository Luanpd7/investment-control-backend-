package database

import (
	"database/sql"
    "os"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	var err error


    DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        return err
    }

    log.Println("2 - sql.Open executado")
    log.Println("3 - DB:", DB)

	if err = DB.Ping(); err != nil {
		return err
	}

	log.Println("✅ Banco conectado!")

	return createTables()
}

func createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS investments (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP NOT NULL,
    emergency NUMERIC(15,2) NOT NULL DEFAULT 0,
    fixed_income NUMERIC(15,2) NOT NULL DEFAULT 0,
    variable_income NUMERIC(15,2) NOT NULL DEFAULT 0,
    contribution NUMERIC(15,2) NOT NULL DEFAULT 0,
    variation NUMERIC(15,2) NOT NULL DEFAULT 0,
    total NUMERIC(15,2) NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
	`

	_, err := DB.Exec(query)
	return err
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
