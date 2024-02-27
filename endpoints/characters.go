// endpoints/characters.go
package endpoints

import (
	"context"
	"fmt"
	"github.com/JanSolo1/sw-api/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

func RegisterCharacterRoutes(router *gin.Engine, db *mongo.Database) {
	router.POST("/characters", func(c *gin.Context) {
		var character models.StarWarsCharacter
		if err := c.ShouldBindJSON(&character); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := db.Collection("starwarscharacter")
		res, err := collection.InsertOne(c, character)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting StarWarsCharacter"})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	router.GET("/characters", func(c *gin.Context) {
		collection := db.Collection("starwarscharacter")
		cursor, err := collection.Find(c, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching StarWarsCharacters"})
			return
		}

		var characters []models.StarWarsCharacter
		if err = cursor.All(c, &characters); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding StarWarsCharacters"})
			return
		}

		c.JSON(http.StatusOK, characters)
	})

	router.PUT("/characters/:id", func(c *gin.Context) {
		var character models.StarWarsCharacter
		if err := c.ShouldBindJSON(&character); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		collection := db.Collection("starwarscharacter")
		res, err := collection.UpdateOne(c, bson.M{"_id": c.Param("id")}, bson.M{"$set": character})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating StarWarsCharacter"})
			return
		}

		c.JSON(http.StatusOK, res)
	})

	router.DELETE("/characters/:id", func(c *gin.Context) {
		collection := db.Collection("starwarscharacter")
		res, err := collection.DeleteOne(c, bson.M{"_id": c.Param("id")})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting StarWarsCharacter"})
			return
		}

		c.JSON(http.StatusOK, res)
	})
}

func GetAllCharacters(c *gin.Context, db *mongo.Database) {
	fmt.Println("Getting collection: starwarscharacter") // Debugging line
	collection := db.Collection("starwarscharacter")

	fmt.Println("Creating context with timeout") // Debugging line
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var characters []models.StarWarsCharacter

	fmt.Println("Finding characters in database") // Debugging line
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Printf("Error finding characters: %v\n", err) // Debugging line
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching characters from database"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var character models.StarWarsCharacter
		fmt.Println("Decoding character data") // Debugging line
		if err := cursor.Decode(&character); err != nil {
			fmt.Printf("Error decoding character: %v\n", err) // Debugging line
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding character data"})
			return
		}
		characters = append(characters, character)
	}

	if err := cursor.Err(); err != nil {
		fmt.Printf("Error with cursor: %v\n", err) // Debugging line
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error with database cursor"})
		return
	}

	fmt.Println("Successfully fetched characters") // Debugging line
	c.JSON(http.StatusOK, characters)
}
