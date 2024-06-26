package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct
type Netflix struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` // use omitempty for prevent defauld id [{ObjectID("000000000000000000000000")] clash
	Movie string `json:"movie" bson:"movie"`
	Watched bool `json:"watched" bson:"watched"`
}