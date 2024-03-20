package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Username  *string            `bson:"username" json:"username"`
	Avatar    *string            `bson:"avatar" json:"avatar"`
	Diamond   int                `bson:"diamond" json:"diamond"`
	Score     int                `bson:"score" json:"score"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
