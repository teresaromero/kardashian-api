package controllers

import (
	"kardashian_api/models"
	"kardashian_api/repository"
	"kardashian_api/utils/http_errors"
)

func Collection(tableName string) (interface{}, *http_errors.HttpError) {
	if tableName == string(models.WikiEpisodes) {
		return repository.GetAllWikiEpisodes()
	} else if tableName == string(models.IMBDEpisodes) {
		return repository.GetAllIMBDEpisodes()
	} else if tableName == string(models.IMBDEpisodeCredits) {
		return repository.GetAllIMBDEpisodeCredits()
	}
	return nil, http_errors.InvalidCollection(tableName)

}

func AvailableCollections(baseURL string) ([]*models.AvailableCollection, *http_errors.HttpError) {
	var rsp []*models.AvailableCollection

	colList, err := repository.GetAvailableCollections()
	if err != nil {
		return nil, err
	}
	for _, collection := range colList {
		rsp = append(rsp, &models.AvailableCollection{Name: collection, Url: baseURL + collection})
	}
	return rsp, nil
}
