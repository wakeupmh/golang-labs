package handlers 

import (
	"database/sql"  
	"log"           
	"net/http"      
	"strconv"       

	"github.com/gin-gonic/gin" 
	"github.com/go-playground/validator/v10"

	"my-books/models" 
)

type BookHandler struct {
	DB *sql.DB 
}

func NewBookHandler(db *sql.DB) *BookHandler {
	return &BookHandler{DB: db}
}

func (bh *BookHandler) GetBooks(c *gin.Context) {
	rows, err := bh.DB.Query("SELECT id, title, author, year FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching books"})
		log.Printf("error fetching books: %v", err)
		return
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error scanning book"})
			log.Printf("error scanning book: %v", err)
			return
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error iterating books"})
		log.Printf("error iterating books: %v", err)
		return
	}

	c.JSON(http.StatusOK, books)
}

func (bh *BookHandler) GetBookByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book ID"})
		return
	}

	var book models.Book
	row := bh.DB.QueryRow("SELECT id, title, author, year FROM books WHERE id = ?", id)
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error fetching book"})
		log.Printf("error fetching book: %v", err)
		return
	}

	c.JSON(http.StatusOK, book)
}

func (bh *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
		log.Printf("error decoding JSON: %v", err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error: " + err.Error()})
		log.Printf("validation error: %v", err)
		return
	}

	result, err := bh.DB.Exec("INSERT INTO books (title, author, year) VALUES (?, ?, ?)", book.Title, book.Author, book.Year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error creating book"})
		log.Printf("error creating book: %v", err)
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting ID of book"})
		log.Printf("error getting LastInsertId: %v", err)
		return
	}
	book.ID = int(id)

	c.JSON(http.StatusCreated, book) 
}

func (bh *BookHandler) UpdateBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do livro inválido"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON: " + err.Error()})
		log.Printf("error decoding JSON: %v", err)
		return
	}

	validate := validator.New()
	if err := validate.Struct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "validation error: " + err.Error()})
		log.Printf("validation error: %v", err)
		return
	}

	result, err := bh.DB.Exec("UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?", book.Title, book.Author, book.Year, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error updating book"})
		log.Printf("error updating book: %v", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error verifying rows affected"})
		log.Printf("error getting RowsAffected: %v", err)
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found for update"})
		return
	}

	book.ID = id 
	c.JSON(http.StatusOK, book)
}

func (bh *BookHandler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do livro inválido"})
		return
	}

	result, err := bh.DB.Exec("DELETE FROM books WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting book"})
		log.Printf("error deleting book: %v", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error verifying rows affected"})
		log.Printf("error getting RowsAffected: %v", err)
		return
	}
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found for deletion"})
		return
	}

	c.Status(http.StatusNoContent) 
}
