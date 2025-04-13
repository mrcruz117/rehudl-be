package main

import (
	"log"
	"net/http"

	"github.com/mrcruz117/rehudl-be/internal/db"
)

func main() {
	db, err := db.InitSQLiteDB("./internal/db/rehudl.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
