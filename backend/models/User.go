package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserName     string             `json:"username" validate:"required,min=2,max=100"`
	Name         string             `json:"name" validate:"required,min=2,max=100"`
	Password     string             `json:"password" validate:"required,min=8,max=100"`
	UserIdentity string             `json:"userid" validate:"required,len=8"`
	Birthday     string             `json:"birthday" validate:"required,datetime=02-01-2006"`
	UserType     string             `json:"usertype" validate:"eq=janitor|eq=collector|eq=backofficer"`
	SignedToken  string             `json:"signedtoken"`
	RefreshToken string             `json:"refreshtoken"`
}
