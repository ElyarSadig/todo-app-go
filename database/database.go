package database

import (
	"database/sql"
	"fmt"
)

func New(dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	return db, nil
}

func CreateTables(db *sql.DB) error {
	_, err := db.Exec(user)
	if err != nil {
		return err
	}
	_, err = db.Exec(todo)
	if err != nil {
		return err
	}
	return nil
}
