package controllers

import (
	"kardashian_api/models"
	"kardashian_api/repository"
	"kardashian_api/utils/http_errors"
)

func GetAllEpisodes(pagination *models.PaginationOpts) ([]models.Episode, *http_errors.HttpError) {
	return repository.GetAllEpisodes(pagination)
}

func GetEpisodeByNumber(n int) (interface{}, *http_errors.HttpError) {
	return repository.GetEpisodeByNumber(n)
}
