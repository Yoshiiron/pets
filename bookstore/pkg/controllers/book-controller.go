package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yoshiiron/bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	c.JSON(http.StatusOK, res)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while parsing",
		})
		return
	}
	bookDetails, _ := models.GetBookById(ID)
	c.JSON(http.StatusOK, gin.H{
		"result": bookDetails,
	})
}

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}
	c.ShouldBindJSON(&CreateBook)
	b := CreateBook.CreateBook()
	c.JSON(http.StatusOK, gin.H{
		"result": b,
	})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("bookId")
	ID, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while parsing",
		})
		return
	}
	book := models.DeleteBook(ID)
	c.JSON(http.StatusOK, gin.H{
		"result": book,
	})
}

func UpdateBook(c *gin.Context) {
	var UpdateBook = &models.Book{}
	c.BindJSON(&UpdateBook)
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error while parsing",
		})
	}
	bookDetails, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	c.JSON(http.StatusOK, gin.H{
		"result": bookDetails,
	})
}
