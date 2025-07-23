package auth

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type User struct {
	ID        primitive.ObjectID `json:"id" binding:"required"`
	Username  string             `json:"username" binding:"required"`
	Email     string             `json:"email" binding:"required,email"`
	Password  string             `json:"password" binding:"required"`
	CreatedAt bson.DateTime      `json:"created_at" binding:"required"`
}

type SigningUpUser struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
