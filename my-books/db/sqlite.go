package db

import (
	"database/sql"
	"log"
	
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() (*sql.DB, error)  {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		author TEXT NOT NULL,
		year INTEGER NOT NULL
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("error creating table: %v", err)
	}
	log.Println("database initialized and table 'books' verified.")

	return db, nil
}
