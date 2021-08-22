package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type IMBDEpisode struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id"`
	ImbdId        string             `json:"imbd_id" bson:"imbd_id"`
	AirDate       primitive.DateTime `json:"air_date" bson:"air_date"`
	RawAirDate    string             `json:"raw_air_date" bson:"raw_air_date"`
	Title         string             `json:"title" bson:"title"`
	Description   string             `json:"description" bson:"description"`
	ImageUrl      string             `json:"image_url" bson:"image_url"`
	ImbdRate      float64            `json:"imbd_rate" bson:"imbd_rate"`
	ImbdRateVotes int                `json:"imbd_rate_votes" bson:"imbd_rate_votes"`
	Season        int                `json:"season" bson:"season"`
	SeasonEpisode int                `json:"episode" bson:"episode"`
}
