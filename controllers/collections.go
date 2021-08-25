package controllers

import (
	"kardashian_api/custom_errors"
	"kardashian_api/models"
	"kardashian_api/repository"
)

func Collection(tableName string) (interface{}, *custom_errors.HttpError) {
	if tableName == string(models.WikiEpisodes) {
		return repository.GetAllWikiEpisodes()
	} else if tableName == string(models.IMBDEpisodes) {
		return repository.GetAllIMBDEpisodes()
	} else if tableName == string(models.IMBDEpisodeCredits) {
		return repository.GetAllIMBDEpisodeCredits()
	}
	return nil, custom_errors.InvalidCollection(tableName)

}

func AvailableCollections(baseURL string) ([]*models.AvailableCollection, *custom_errors.HttpError) {
	return repository.GetAvailableCollections(baseURL)
}
