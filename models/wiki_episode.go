package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const WikiEpisodes Collection = "wiki_episodes"

type WikiEpisode struct {
	Id            primitive.ObjectID `json:"_id" bson:"_id"`
	AirDate       primitive.DateTime `json:"air_date" bson:"air_date"`
	Title         string             `json:"title" bson:"title"`
	SeasonEpisode int                `json:"episode" bson:"episode"`
	Episode       int                `json:"episode_overall" bson:"episode_overall"`
	Special       bool               `json:"special_episode" bson:"special_episode"`
	Season        string             `json:"season,omitempty" bson:"season,omitempty"` // Optional
	Viewers       float64            `json:"us_viewers" bson:"us_viewers"`             // Optional
}
