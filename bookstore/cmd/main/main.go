package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yoshiiron/bookstore/pkg/routes"
)

func main() {

	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("error loading envs from .env")
	}

	api_port := os.Getenv("API_PORT")

	r := gin.Default()

	routes.RegisterBookStoreRoutes(r)

	log.Fatal(r.Run(api_port))
}
