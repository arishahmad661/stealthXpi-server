package auth

import (
	"context"
	"github.com/arishahmad661/stealth_x_pi/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

func RegisterUser(signingUpUser SigningUpUser) (User, error) {
	var user User
	collection := config.DB.Collection("users")

	user.ID = primitive.NewObjectID()
	user.Username = signingUpUser.Username
	user.Email = signingUpUser.Email
	user.Password = signingUpUser.Password
	user.CreatedAt = bson.DateTime(time.Now().Unix())

	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return User{}, err
	}

	return User{}, nil
}
