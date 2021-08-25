package models

const Episodes Collection = "episodes"

type Episode struct {
	Id          string `json:"imbd_id" bson:"imbd_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`

	AirDate int `json:"air_date" bson:"air_date"`

	Season        int `json:"season" bson:"season"`
	Episode       int `json:"episode" bson:"episode"`
	EpisodeNumber int `json:"episode_overall" bson:"episode_overall"`

	Special bool    `json:"special_episode" bson:"special_episode"`
	Viewers float64 `json:"us_viewers" bson:"us_viewers"`

	Rate  float64 `json:"imbd_rate" bson:"imbd_rate"`
	Votes int     `json:"imbd_rate_votes" bson:"imbd_rate_votes"`
}
