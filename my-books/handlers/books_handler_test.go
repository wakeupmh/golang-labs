package handlers

import (
	"bytes"        
	"database/sql"  
	"encoding/json" 
	"net/http"      
	"net/http/httptest" 
	"testing"       
	"fmt"

	_ "github.com/mattn/go-sqlite3" 
	"my-books/models"    
)

func setupTestDB(t *testing.T) (*sql.DB, func()) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("error while opening database: %v", err)
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
		db.Close()
		t.Fatalf("error while creating table books: %v", err)
	}

	return db, func() {
		db.Close()
	}
}

func TestGetBooks(t *testing.T) {
	db, teardown := setupTestDB(t) 
	defer teardown()               

	bh := NewBookHandler(db) 

	req, _ := http.NewRequest("GET", "/books", nil) 
	rr := httptest.NewRecorder()
	bh.GetBooks(rr, req)                            

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: expected %v, got %v",
			http.StatusOK, status)
	}

	expected := "[]\n" 
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: expected %v, got %v",
			expected, rr.Body.String())
	}

	db.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", "Livro A", "Autor A", 2000)
	db.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", "Livro B", "Autor B", 2010)

	req, _ = http.NewRequest("GET", "/books", nil)
	rr = httptest.NewRecorder()
	bh.GetBooks(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: expected %v, got %v",
			http.StatusOK, status)
	}

	var books []models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatalf("error while decoding JSON response: %v", err)
	}
	if len(books) != 2 {
		t.Errorf("expected 2 books, got %d", len(books))
	}
	if books[0].Title != "Livro A" || books[0].Author != "Autor A" {
		t.Errorf("unexpected book data: %v", books[0])
	}
}

func TestCreateBook(t *testing.T) {
	db, teardown := setupTestDB(t)
	defer teardown()
	bh := NewBookHandler(db)

	bookJSON := []byte(`{"title": "Novo Livro", "author": "Novo Autor", "year": 2023}`)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(bookJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	bh.CreateBook(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: expected %v, got %v",
			http.StatusCreated, status)
	}

	var createdBook models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &createdBook)
	if err != nil {
		t.Fatalf("error while decoding JSON response: %v", err)
	}

	if createdBook.ID == 0 || createdBook.Title != "Novo Livro" {
		t.Errorf("unexpected book data: %v", createdBook)
	}
	invalidJSON := []byte(`{"title": "Livro Mal Formado", "author": "Autor Mal Formado", "year": "dois mil e vinte e quatro"}`)
	req, _ = http.NewRequest("POST", "/books", bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	bh.CreateBook(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid JSON: expected %v, got %v",
			http.StatusBadRequest, status)
	}
}

func TestGetBookByID(t *testing.T) {
	db, teardown := setupTestDB(t)
	defer teardown()
	bh := NewBookHandler(db)

	res, _ := db.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", "Livro Teste ID", "Autor Teste", 2020)
	insertedID, _ := res.LastInsertId()

	req, _ := http.NewRequest("GET", fmt.Sprintf("/books/%d", insertedID), nil)
	rr := httptest.NewRecorder()
	bh.GetBookByID(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: expected %v, got %v",
			http.StatusOK, status)
	}

	var book models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatalf("error while decoding JSON response: %v", err)
	}
	if book.ID != int(insertedID) || book.Title != "Livro Teste ID" {
		t.Errorf("unexpected book data: %v", book)
	}

	req, _ = http.NewRequest("GET", "/books/9999", nil)
	rr = httptest.NewRecorder()
	bh.GetBookByID(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for non-existent book: expected %v, got %v",
			http.StatusNotFound, status)
	}

	req, _ = http.NewRequest("GET", "/books/abc", nil)
	rr = httptest.NewRecorder()
	bh.GetBookByID(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid ID: expected %v, got %v",
			http.StatusBadRequest, status)
	}
}

func TestUpdateBook(t *testing.T) {
	db, teardown := setupTestDB(t)
	defer teardown()
	bh := NewBookHandler(db)
	res, _ := db.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", "Livro Antigo", "Autor Antigo", 2000)
	insertedID, _ := res.LastInsertId()

	updatedBookJSON := []byte(`{"title": "Livro Atualizado", "author": "Autor Atualizado", "year": 2024}`)
	req, _ := http.NewRequest("PUT", fmt.Sprintf("/books/%d", insertedID), bytes.NewBuffer(updatedBookJSON))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	bh.UpdateBook(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: expected %v, got %v",
			http.StatusOK, status)
	}

	var updatedBook models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &updatedBook)
	if err != nil {
		t.Fatalf("error while decoding JSON response: %v", err)
	}
	if updatedBook.Title != "Livro Atualizado" || updatedBook.Author != "Autor Atualizado" {
		t.Errorf("unexpected book data: %v", updatedBook)
	}

	nonExistentID := 9999
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/books/%d", nonExistentID), bytes.NewBuffer(updatedBookJSON))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	bh.UpdateBook(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for non-existent book: expected %v, got %v",
			http.StatusNotFound, status)
	}

	invalidJSON := []byte(`{"title": "Livro Inválido", "year": "ano inválido"}`)
	req, _ = http.NewRequest("PUT", fmt.Sprintf("/books/%d", insertedID), bytes.NewBuffer(invalidJSON))
	req.Header.Set("Content-Type", "application/json")
	rr = httptest.NewRecorder()
	bh.UpdateBook(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid JSON: expected %v, got %v",
			http.StatusBadRequest, status)
	}
}

func TestDeleteBook(t *testing.T) {
	db, teardown := setupTestDB(t)
	defer teardown()
	bh := NewBookHandler(db)

	res, _ := db.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", "Livro para Deletar", "Autor", 2000)
	insertedID, _ := res.LastInsertId()

	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/books/%d", insertedID), nil)
	rr := httptest.NewRecorder()
	bh.DeleteBook(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: expected %v, got %v",
			http.StatusNoContent, status)
	}

	var count int
	db.QueryRow("SELECT COUNT(*) FROM books WHERE id = ?", insertedID).Scan(&count)
	if count != 0 {
		t.Errorf("book was not deleted from the database.")
	}
	nonExistentID := 9999
	req, _ = http.NewRequest("DELETE", fmt.Sprintf("/books/%d", nonExistentID), nil)
	rr = httptest.NewRecorder()
	bh.DeleteBook(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for non-existent book: expected %v, got %v",
			http.StatusNotFound, status)
	}
	req, _ = http.NewRequest("DELETE", "/books/abc", nil)
	rr = httptest.NewRecorder()
	bh.DeleteBook(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for invalid ID: expected %v, got %v",
			http.StatusBadRequest, status)
	}
}
