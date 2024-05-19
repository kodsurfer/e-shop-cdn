package dtos

import "go.mongodb.org/mongo-driver/bson/primitive"

type SubscribePayloadDto struct {
	Topic string `json:"topic"`
}

type UnsubscribePayloadDto struct {
	TopicId primitive.ObjectID `json:"topic_id"`
}
