package main

import (
	"net/http"
	"submission-project-enigma-laundry/config"

	"github.com/gin-gonic/gin"
)

type Customers struct {
	Id          int    `json:"id"`
	Name       string `json:"name"`
	PhoneNumber      string `json:"phoneNumber"`
	Address string `json:"address"`
}

var db = config.ConnectDB()

func main() {
	// Tulis kode kamu disini
	router := gin.Default()
	// router.Use(LoggerMiddleware)

	// apiGroup := router.Group("/api")
	// {
	// 		booksGroup := apiGroup.Group("/customers")
	// 		{
	// 			booksGroup.GET("/", getAllBook)
	// 			booksGroup.POST("/", create_Customer)
	// 			booksGroup.GET("/:id", getBookById)
	// 			booksGroup.PUT("/:id", updateBookById)
	// 		}
	// }
	router.POST("/customers", create_Customer)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func create_Customer(c *gin.Context) {
	var newBook Customers
	err := c.ShouldBind(&newBook)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := "INSERT INTO Customers (name, phoneNumber, address) VALUES ($1, $2, $3) RETURNING id"

	var bookId int
	err = db.QueryRow(query, newBook.Name, newBook.PhoneNumber, newBook.Address).Scan(&bookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed to create book"})
		return
	}

	newBook.Id = bookId
	c.JSON(http.StatusCreated, newBook)
}
