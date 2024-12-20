package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Email     string             `bson:"email" json:"email"`
	Subject   string             `bson:"subject" json:"subject"`
	Message   string             `bson:"message" json:"message"`
	CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
}
