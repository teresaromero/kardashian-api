package controllers

import (
	"kardashian_api/custom_errors"
	"kardashian_api/models"
	"kardashian_api/repository"
)

func GetAllEpisodes(pagination *models.PaginationOpts) ([]models.Episode, *custom_errors.HttpError) {
	return repository.GetAllEpisodes(pagination)
}

func GetEpisodeByNumber(n int) (interface{}, *custom_errors.HttpError) {
	return repository.GetEpisodeByNumber(n)
}
