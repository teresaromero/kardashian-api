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

	cursor, err_find := collection.Find(ctx, bson.D{})
	if err_find != nil {
		panic(err_find)
		}
		defer cursor.Close(ctx)

		// TODO: parse each collection model, using Episode as placeholder for now
	err_cursor := cursor.All(ctx, &coll_items)
	if err_cursor != nil {
		panic((err_cursor))
		}
		c.JSON(200, coll_items)
	}

	c.JSON(400, gin.H{"error": "Collection not valid"})

}
