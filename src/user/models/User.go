package models

import (
	"time"

	"github.com/shipu/artifact"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var UserCollection artifact.MongoCollection

type User struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Task      string             `json:"task" bson:"task"`
	Status    string             `json:"status" bson:"status"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func UserSetup() {
	UserCollection = artifact.Mongo.Collection("users")
}
