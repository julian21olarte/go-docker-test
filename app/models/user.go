package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User model
type User struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name  string             `json:"name" validate:"required"`
	Age   int                `json:"age" validate:"required"`
	Phone string             `json:"phone" validate:"required"`
	Email string             `json:"email" validate:"required,email"`
}
