package controllers

import (
	"kardashian_api/database"
	"kardashian_api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func Collection(tableName string) (items []models.IMBDEpisode, err error) {
	collection := database.Use(tableName)

	ctx, cancel := database.Context()
	defer cancel()

	cursor, err_find := collection.Find(ctx, bson.D{})
	if err_find != nil {
		return nil, err_find
	}
	defer cursor.Close(ctx)

	// TODO: parse each collection model, using Episode as placeholder for now
	var coll_items []models.IMBDEpisode
	err_cursor := cursor.All(ctx, &coll_items)
	if err_cursor != nil {
		return nil, err_cursor

	}
	return coll_items, nil

}

func AvailableCollections(baseURL string) ([]*models.AvailableCollection, error) {

	cls, err := database.Collections()
	if err != nil {
		return nil, err
	}

	var rsp []*models.AvailableCollection
	baseUrl := baseURL
	for _, col := range cls {
		rsp = append(rsp, &models.AvailableCollection{Name: col, Url: baseUrl + col})
	}

	return rsp, nil

}
