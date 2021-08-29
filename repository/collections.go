package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"kardashian_api/database"
	"kardashian_api/models"
	"kardashian_api/utils/http_errors"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAvailableCollections() ([]string, *http_errors.HttpError) {

	lcn, err := database.DB.ListCollectionNames(context.TODO(), bson.M{})
	if err != nil {
		return nil, http_errors.InternalServerError(err)
	}

	return lcn, nil
}

func GetAllCollection(items interface{}, collection string) *http_errors.HttpError {

	ctx := context.TODO()
	cursor, errFind := database.DB.Collection(collection).Find(ctx, bson.M{})
	if errFind != nil {
		return http_errors.InternalServerError(errFind)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
		}
	}(cursor, ctx)

	errAll := cursor.All(ctx, &items)
	if errAll != nil {
		return http_errors.InternalServerError(errAll)
	}

	return nil
}

func GetAllIMBDEpisodes() (items []models.IMBDEpisode, err *http_errors.HttpError) {
	err = GetAllCollection(items, string(models.IMBDEpisodes))
	if err != nil {
		return nil, err
	}
	return items, nil
}

func GetAllWikiEpisodes() (items []models.WikiEpisode, err *http_errors.HttpError) {
	err = GetAllCollection(items, string(models.WikiEpisodes))
	if err != nil {
		return nil, err
	}
	return items, nil

}

func GetAllIMBDEpisodeCredits() (items []models.IMBDEpisodeCredit, err *http_errors.HttpError) {
	err = GetAllCollection(items, string(models.IMBDEpisodeCredits))
	if err != nil {
		return nil, err
	}
	return items, nil
}
