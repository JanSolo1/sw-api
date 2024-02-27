// endpoints/characters.go
package endpoints

import (
	"github.com/JanSolo1/sw-api/models"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func createCharacter(c *gin.Context, collection *mongo.Collection) {
	var character models.StarWarsCharacter
	if err := c.ShouldBindJSON(&character); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"body":  c.Request.Body,
		}).Error("Error binding character JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding character JSON: " + err.Error()})
		return
	}

	res, err := collection.InsertOne(c, character)
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"character": character,
		}).Error("Error inserting StarWarsCharacter")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting StarWarsCharacter: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func getCharacters(c *gin.Context, collection *mongo.Collection) {
	cursor, err := collection.Find(c, bson.M{})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error fetching StarWarsCharacters")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching StarWarsCharacters: " + err.Error()})
		return
	}

	var characters []models.StarWarsCharacter
	if err = cursor.All(c, &characters); err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error decoding StarWarsCharacters")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding StarWarsCharacters: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, characters)
}

func updateCharacter(c *gin.Context, collection *mongo.Collection) {
	var character models.StarWarsCharacter
	if err := c.ShouldBindJSON(&character); err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"body":  c.Request.Body,
		}).Error("Error binding character JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding character JSON: " + err.Error()})
		return
	}

	res, err := collection.UpdateOne(c, bson.M{"_id": c.Param("id")}, bson.M{"$set": character})
	if err != nil {
		log.WithFields(log.Fields{
			"error":     err,
			"character": character,
		}).Error("Error updating StarWarsCharacter")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating StarWarsCharacter: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func deleteCharacter(c *gin.Context, collection *mongo.Collection) {
	res, err := collection.DeleteOne(c, bson.M{"_id": c.Param("id")})
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error deleting StarWarsCharacter")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting StarWarsCharacter: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, res)
}

func RegisterCharacterRoutes(router *gin.Engine, db *mongo.Database) {
	collection := db.Collection("starwarscharacter")

	router.POST("/characters", func(c *gin.Context) {
		createCharacter(c, collection)
	})

	router.GET("/characters", func(c *gin.Context) {
		getCharacters(c, collection)
	})

	router.PUT("/characters/:id", func(c *gin.Context) {
		updateCharacter(c, collection)
	})

	router.DELETE("/characters/:id", func(c *gin.Context) {
		deleteCharacter(c, collection)
	})
}
