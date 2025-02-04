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
