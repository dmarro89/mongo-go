package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `json:"id"`
	Name    string             `json:"name" validate:"required"`
	Surname string             `json:"surname" validate:"required"`
	Email   string             `json:"email,omitempty" validate:"required"`
}
