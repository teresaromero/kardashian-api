package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"kardashian_api/database"
	"kardashian_api/models"
	"kardashian_api/utils/http_errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetAllEpisodes(p *models.PaginationOpts) ([]models.Episode, *http_errors.HttpError) {

	opts := options.Find()
	opts.SetLimit(int64(p.Limit))
	opts.SetSkip(int64(p.Skip))
	opts.SetSort(bson.D{{Key: "episode_overall", Value: 1}})

	filter := bson.D{}

	ctx, cancel := database.Context()
	defer cancel()

	cursor, err := database.DB.Collection(string(models.Episodes)).Find(ctx, filter, opts)
	if err != nil {
		return nil, http_errors.InternalServerError(err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
		}
	}(cursor, ctx)

	var episodes []models.Episode
	err = cursor.All(ctx, &episodes)
	if err != nil {
		return nil, http_errors.InternalServerError(err)
	}

	return episodes, nil
}

func GetEpisodeByNumber(num int) (interface{}, *http_errors.HttpError) {

	ctx, cancel := database.Context()
	defer cancel()

	result := database.DB.Collection(string(models.Episodes)).FindOne(ctx, bson.D{primitive.E{Key: "episode_overall", Value: num}})

	var episode models.Episode
	err := result.Decode(&episode)
	if err != nil {
		return nil, http_errors.BadRequest(err)
	}
	return episode, nil
}
