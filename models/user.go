package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id      primitive.ObjectID `json:"id"`
	Name    string             `form:"name" json:"name" binding:"required"`
	Surname string             `form:"username" json:"surname" binding:"required"`
	Email   string             `form:"email" json:"email" binding:"required"`
}
