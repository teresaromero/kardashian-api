package controllers

import (
	"context"
	"kardashian_api/database"
	"kardashian_api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ParseCollection(cursor *mongo.Cursor, tableName string, ctx context.Context) (interface{}, error) {
	var imbd_episode_items []models.IMBDEpisode
	var imbd_episode_credit_items []models.IMBDEpisodeCredit

	var wiki_episode_items []models.WikiEpisode
	var err error
	var result interface{}

	if tableName == string(models.WikiEpisodes) {
		err = cursor.All(ctx, &wiki_episode_items)
		result = wiki_episode_items
	} else if tableName == string(models.IMBDEpisodes) {
		err = cursor.All(ctx, &imbd_episode_items)
		result = imbd_episode_items
	} else if tableName == string(models.IMBDEpisodeCredits) {
		err = cursor.All(ctx, &imbd_episode_credit_items)
		result = imbd_episode_credit_items
	}
	return result, err
}

func Collection(tableName string) (items interface{}, err error) {
	collection := database.Use(tableName)

	ctx, cancel := database.Context()
	defer cancel()

	cursor, err_find := collection.Find(ctx, bson.D{})
	if err_find != nil {
		return nil, err_find
	}
	defer cursor.Close(ctx)

	result, err := ParseCollection(cursor, tableName, ctx)
	if err != nil {
		return nil, err
	}

	return result, nil

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
