// web service to perform CRUD on set of books
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var Books = []book{
	{ID: "1", Title: "Blue Train", Author: "John Coltrane"},
	{ID: "2", Title: "Jeru", Author: "Gerry Mulligan"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Author: "Sarah Vaughan"},
}

func main() {
	router := gin.Default()
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookById)
	router.POST("/books", postBooks)
	router.PUT("/books/:id", updateBook)
	router.DELETE("/books/:id", deleteBook)
	router.Run("localhost:9000")
}

// GET
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Books)
}

// GET BY ID
func getBookById(c *gin.Context) {
	id := c.Param("id")

	for _, b := range Books {
		if b.ID == id {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// POST
func postBooks(c *gin.Context) {
	var newBook book
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	Books = append(Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// UPDATE
func updateBook(c *gin.Context) {
	id := c.Param("id")
	var newBook book

	// get body in newBook
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// iterate the slice and update
	for index, b := range Books {
		if b.ID == id {
			if newBook.Author != b.Author && newBook.Author != "" {
				Books[index].Author = newBook.Author
			}
			if newBook.Title != b.Title && newBook.Title != "" {
				Books[index].Title = newBook.Title
			}

		}
	}

	//send response
	c.IndentedJSON(http.StatusOK, gin.H{"message": "book updated"})
}

// DELETE
func deleteBook(c *gin.Context) {
	id := c.Param("id")

	// iterate and delete the element
	for index, book := range Books {
		if book.ID == id {
			Books = append(Books[:index], Books[index+1:]...)

		}
	}

	for _, book := range Books {
		fmt.Println(book.ID, book.Title, book.Author)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "success"})
}
