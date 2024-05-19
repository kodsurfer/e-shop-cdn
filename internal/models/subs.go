package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SubModel struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"` // from Auth service
	Topic     string             `json:"topic" bson:"topic"`     // /bucket/pathA/pathB/filename
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at"`
}

type UserSubModel struct {
	Topic string
}

type PaginatedSubs struct {
	Total int64      `json:"total"`
	Data  []SubModel `json:"data"`
}

type ISubsFilter struct {
	UserId string `json:"user_id" bson:"user_id"`
}
