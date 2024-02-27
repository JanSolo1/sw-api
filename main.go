// main.go
package main

import (
	"fmt"
	"github.com/JanSolo1/sw-api/endpoints"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func setupGin(db *mongo.Database) *gin.Engine {
	// Set up Gin
	r := gin.Default()

	// Existing route
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})

	// New route for fetching all characters
	//r.GET("/characters", func(c *gin.Context) {
	//	endpoints.GetAllCharacters(c, db)
	//})

	endpoints.RegisterCharacterRoutes(r, db)

	return r
}

func main() {
	fmt.Println("Star Wars API")
	PORT := "8080"

	r := setupGin(db)

	// Use the MongoDB client and Gin engine here...
	// Run the server
	err := r.Run(":" + PORT)
	if err != nil {
		log.Fatal(err)
		return
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
