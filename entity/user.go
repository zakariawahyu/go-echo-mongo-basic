package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty"`
	FirstName string             `json:"first_name,omitempty" validate:"required"`
	LastName  string             `json:"last_name,omitempty" validate:"required"`
	Username  string             `json:"username,omitempty" validate:"required"`
	Location  string             `json:"location,omitempty" validate:"required"`
	Title     string             `json:"title,omitempty" validate:"required"`
}
