package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Calendar struct {
	Id    primitive.ObjectID `bson:"_id"`
	Route string             `json:"route"`
	Time  string             `json:"time" validate:"required,datetime=15:04 02-01-2006"`
}
