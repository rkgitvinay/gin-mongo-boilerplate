package models

import (
	"time"

	"github.com/shipu/artifact"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var TodoCollection artifact.MongoCollection

type Todo struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Task      string             `json:"task" bson:"task"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func TodoSetup() {
	TodoCollection = artifact.Mongo.Collection("users")
}
