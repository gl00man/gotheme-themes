package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Name: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Quantity: 10},
	{ID: "2", Name: "Cloud Native Go", Author: "M.-L. Reimer, K. Köhler", Quantity: 5},
	{ID: "3", Name: "To kill a mockingbird", Author: "Harper Lee", Quantity: 3},
	{ID: "4", Name: "The Lord of the Rings", Author: "J. R. R. Tolkien", Quantity: 1},
	{ID: "5", Name: "The Da Vinci Code", Author: "Dan Brown", Quantity: 2},
	{ID: "6", Name: "The Catcher in the Rye", Author: "J. D. Salinger", Quantity: 4},
	{ID: "7", Name: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 6},
	{ID: "8", Name: "The Hobbit", Author: "J. R. R. Tolkien", Quantity: 7},
	{ID: "9", Name: "War and Peace", Author: "Leo Tolstoy", Quantity: 8},
	{ID: "10", Name: "The Hunger Games", Author: "Suzanne Collins", Quantity: 9},
	{ID: "11", Name: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 10},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)

}

func createBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)

}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("Book not found")
}
func bookById(c *gin.Context) {
	id := c.Param("id")

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.POST("/books", createBook)
	router.GET("/books/:id", bookById)
	router.Run("localhost:8083")

}