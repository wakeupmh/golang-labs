package main // Define o pacote principal da aplicação.

import (
	"log"

	"github.com/gin-gonic/gin"

	"my-books/db"
	"my-books/handlers"
)

func main() {
	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Erro fatal ao inicializar o banco de dados: %v", err)
	}
	defer db.Close()

	bookHandler := handlers.NewBookHandler(db)

	router := gin.Default()

	booksGroup := router.Group("/books")
	{
		booksGroup.GET("/", bookHandler.GetBooks)
		booksGroup.POST("/", bookHandler.CreateBook)

		booksGroup.GET("/:id", bookHandler.GetBookByID)
		booksGroup.PUT("/:id", bookHandler.UpdateBook)
		booksGroup.DELETE("/:id", bookHandler.DeleteBook)
	}

	log.Println("server started at :8080")
	log.Fatal(router.Run(":8080"))
}
