package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Calendar struct {
	Id     primitive.ObjectID `bson:"_id"`
	Action string             `json:"action" validate:"required,eq=create||eq=modify||eq=view"`
	Route  string             `json:"route"`
	Time   string             `json:"time" validate:"required,datetime=15:04 02-01-2006"`
}
