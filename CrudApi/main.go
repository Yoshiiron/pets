package main

import (
	"crud-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.InitRoutes(r)

	r.Run(":8080")
}
