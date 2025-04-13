package db

import (
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	db, err := InitSQLiteDB("./rehudl.db")
	if err != nil {
		t.Fatalf("Failed to initialize SQLite database: %v", err)
	}
	defer db.Close()

	// Check if the database file exists
	if _, err := os.Stat("rehudl.db"); os.IsNotExist(err) {
		t.Errorf("Database file does not exist")
	}
}
