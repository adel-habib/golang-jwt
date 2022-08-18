package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	First_name     *string            `json:"first_name"`
	Last_name      *string            `json:"last_name"`
	Email          *string            `json:"email"`
	Password       *string            `json:"password"`
	Token          *string            `json:"token"`
	User_type      *string            `json:"user_type"`
	Referesh_token *string            `json:"referesh_token"`
	Created_at     time.Time          `json:"created_at"`
	Updated_at     time.Time          `json:"updated_at"`
	User_id        string             `json:"user_id"`
}
