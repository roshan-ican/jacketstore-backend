package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FullName     string             `bson:"full_name" json:"full_name"`
	Email        string             `bson:"email" json:"email"`
	PasswordHash string             `bson:"password_hash" json:"-"`
	Phone        string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Address      string             `bson:"address,omitempty" json:"address,omitempty"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at"`
}
