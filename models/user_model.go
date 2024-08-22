package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	// Id       primitive.ObjectID `json:"id,omitempty"`
	// Name     string             `json:"name,omitempty" validate:"required"`
	// Location string             `json:"location,omitempty" validate:"required"`
	// Title    string             `json:"title,omitempty" validate:"required"`
	Id primitive.ObjectID `json:"id,omitempty"`
	Titulo string `json:"titulo,omitempty" validate:"required"`
	Ano int `json:"ano,omitempty" validate:"required"`
	Diretor string `json:"diretor,omitempty" validate:"required"`
}
