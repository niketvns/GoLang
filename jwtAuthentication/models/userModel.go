package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required, min=2, max=100"`
	Last_name     *string            `json:"last_Name" validate:"required, min=2, max=100"`
	Password      *string            `json:"password" validation:"required"`
	Email         *string            `json:"email" validation:"email"`
	Phone         *string            `json:"phone" validation:"required"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validation:"required eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
