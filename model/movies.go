package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID      primitive.ObjectID `json:"id"`
	Titulo  string             `json:"titulo"`
	Ano     string             `json:"ano"`
	Diretor string             `json:"diretor"`
}
