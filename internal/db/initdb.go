package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitSQLiteDB(path string) (*sql.DB, error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	// Enable WAL mode
	if _, err := db.Exec("PRAGMA journal_mode=WAL;"); err != nil {
		return nil, err
	}

	// Enable foreign keys
	if _, err := db.Exec("PRAGMA foreign_keys=ON; "); err != nil {
		return nil, err
	}

	// check if foreign keys are enabled
	var foreignKeysEnabled bool
	if err := db.QueryRow("PRAGMA foreign_keys;").Scan(&foreignKeysEnabled); err != nil {
		return nil, err
	}

	fmt.Println("Foreign keys enabled:", foreignKeysEnabled)

	return db, nil
}
