package models

type Episode struct {
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
}
