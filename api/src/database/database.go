package database

import (
	"database/sql"
	"social-car/src/config"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 driver
)

// Connect opens a database and return it
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", config.StringConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil

}
