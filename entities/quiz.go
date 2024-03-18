package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Quiz struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Question    string             `bson:"question" json:"question"`
	AnswerTrue  string             `bson:"answer_true" json:"answer_true"`
	AnswerFalse []string           `bson:"answer_false" json:"answer_false"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}