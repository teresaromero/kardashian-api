package controllers

import (
	"kardashian_api/database"
	"kardashian_api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Collection(c *gin.Context) {
	coll_param := c.Param("collection")

	if database.ValidCollection(coll_param) {
		collection := database.Use(coll_param)

		ctx, cancel := database.Context()
		defer cancel()

		cursor, err := collection.Find(ctx, bson.D{})
		if err != nil {
			panic(err)
		}
		defer cursor.Close(ctx)

		// TODO: parse each collection model, using Episode as placeholder for now
		var coll_items []models.Episode
		if err = cursor.All(ctx, &coll_items); err != nil {
			panic((err))
		}
		c.JSON(200, coll_items)
	}

	c.JSON(400, gin.H{"error": "Collection not valid"})

}
