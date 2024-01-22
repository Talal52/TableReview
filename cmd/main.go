package main

import (
	"fmt"
	"log"
	"tablereviews/database"
	"tablereviews/routes"
	"github.com/gin-gonic/gin"


)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal("Could not initialize the database:", err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, db)

	fmt.Println("Starting the server on :4000...")
	if err := router.Run(":4000"); err != nil {
		log.Fatal("Failed to start the server:", err)
	}
}
