package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/yoshiiron/bookstore/pkg/routes"
)

func main() {
	r := gin.Default()
	routes.RegisterBookStoreRoutes(r)

	log.Fatal(r.Run(":8080"))
}
