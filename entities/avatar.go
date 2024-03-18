package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Avatar struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Image     string             `bson:"image" json:"image"`
	Price     string             `bson:"price" json:"price"`
	IsLocked  bool               `bson:"isLocked" json:"isLocked"`
	Eqquiped  bool               `bson:"eqquiped" json:"eqquiped"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}
