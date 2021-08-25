package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const IMBDEpisodeCredits Collection = "imbd_episode_credits"

type IMBDEpisodeCredit struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	ImbdId   string             `json:"imbd_id" bson:"imbd_id"`
	Credit   string             `json:"credit" bson:"credit"`
	Category string             `json:"category" bson:"category"`
	Name     string             `json:"name" bson:"name"`
}
