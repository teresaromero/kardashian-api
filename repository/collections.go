package repository

import (
	"kardashian_api/custom_errors"
	"kardashian_api/database"
	"kardashian_api/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAvailableCollections(baseURL string) ([]*models.AvailableCollection, *custom_errors.HttpError) {

	cls, err := database.Collections()
	if err != nil {
		return nil, custom_errors.InternalServerError(err)
	}

	var rsp []*models.AvailableCollection
	baseUrl := baseURL
	for _, col := range cls {
		rsp = append(rsp, &models.AvailableCollection{Name: col, Url: baseUrl + col})
	}

	return rsp, nil
}

func GetAllCollection(items interface{}, collection string) (interface{}, *custom_errors.HttpError) {

	ctx, cancel := database.Context()
	defer cancel()

	cursor, err := database.Use(collection).Find(ctx, bson.D{})

	if err != nil {
		return nil, custom_errors.InternalServerError(err)
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &items)
	if err != nil {
		return nil, custom_errors.InternalServerError(err)
	}

	return items, nil
}

func GetAllIMBDEpisodes() (interface{}, *custom_errors.HttpError) {
	var items []models.IMBDEpisode
	return GetAllCollection(items, string(models.IMBDEpisodes))

}

func GetAllWikiEpisodes() (interface{}, *custom_errors.HttpError) {
	var items []models.WikiEpisode
	return GetAllCollection(items, string(models.WikiEpisodes))

}

func GetAllIMBDEpisodeCredits() (interface{}, *custom_errors.HttpError) {
	var items []models.IMBDEpisodeCredit
	return GetAllCollection(items, string(models.IMBDEpisodeCredits))

}
