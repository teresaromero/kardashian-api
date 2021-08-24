package controllers

import (
	"kardashian_api/database"
	"kardashian_api/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func Collection(c *gin.Context) {
	collection := database.Use(c.Param("collection"))

	ctx, cancel := database.Context()
	defer cancel()

	cursor, err_find := collection.Find(ctx, bson.D{})
	if err_find != nil {
		panic(err_find)
	}
	defer cursor.Close(ctx)

	// TODO: parse each collection model, using Episode as placeholder for now
	var coll_items []models.IMBDEpisode
	err_cursor := cursor.All(ctx, &coll_items)
	if err_cursor != nil {
		panic((err_cursor))
	}
	c.JSON(http.StatusOK, gin.H{"response": coll_items})

}

func AvailableCollections(c *gin.Context) {

	cls, err := database.Collections()
	if err != nil {
		panic(err)
	}

	var rsp []*models.AvailableCollection
	baseUrl := c.Request.URL.String()
	for _, col := range cls {
		rsp = append(rsp, &models.AvailableCollection{Name: col, Url: baseUrl + col})
	}

	c.JSON(http.StatusOK, rsp)

}
