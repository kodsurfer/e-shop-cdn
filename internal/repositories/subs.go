package repositories

import (
	"context"
	"errors"
	mongo2 "github.com/WildEgor/e-shop-cdn/internal/db/mongo"
	"github.com/WildEgor/e-shop-cdn/internal/dtos"
	"github.com/WildEgor/e-shop-cdn/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ISubsRepository interface {
	SubscribeToTopic(uid, topic string) (*models.SubModel, error)
	UnsubscribeFromTopicById(id string)
	FindTopicById(id string) (*models.SubModel, error)
	PaginateUserTopics(filter *models.ISubsFilter, opts *dtos.PaginationOpts) (*models.PaginatedSubs, error)
	FindUserTopics(uid string) ([]string, error)
}

var _ ISubsRepository = (*SubsRepository)(nil)

type SubsRepository struct {
	coll *mongo.Collection
}

func NewSubsRepository(
	db *mongo2.Connection,
) *SubsRepository {

	coll := db.Db().Collection(models.CollectionSubs)

	return &SubsRepository{
		coll,
	}
}

func (s *SubsRepository) FindTopicById(id string) (*models.SubModel, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{Key: "_id", Value: _id}}

	var model *models.SubModel
	if err := s.coll.FindOne(context.TODO(), filter).Decode(&model); err != nil {
		return nil, err
	}

	return model, nil
}

func (s *SubsRepository) SubscribeToTopic(uid, topic string) (*models.SubModel, error) {
	userId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, err
	}

	sub := &models.SubModel{
		UserID:    userId,
		Topic:     topic,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	_, err = s.coll.InsertOne(context.TODO(), sub)
	if err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *SubsRepository) UnsubscribeFromTopicById(id string) {
	_id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.D{{Key: "_id", Value: _id}}
	s.coll.DeleteOne(context.TODO(), filter)
}

func (s *SubsRepository) PaginateUserTopics(filter *models.ISubsFilter, opts *dtos.PaginationOpts) (*models.PaginatedSubs, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result := &models.PaginatedSubs{
		Data: make([]models.SubModel, 0),
	}

	l := opts.Limit
	skip := opts.Page*opts.Limit - opts.Limit

	userId, err := primitive.ObjectIDFromHex(filter.UserId)
	if err != nil {
		return nil, err
	}

	curr, err := s.coll.Find(ctx, bson.D{{
		Key: "user_id", Value: userId,
	}}, &options.FindOptions{Limit: &l, Skip: &skip})
	if err != nil {
		return result, err
	}
	defer curr.Close(ctx)

	count, err := s.coll.CountDocuments(ctx, filter)
	if err != nil {
		return result, err
	}

	result.Total = count

	for curr.Next(ctx) {
		var el models.SubModel
		curr.Decode(&el)

		result.Data = append(result.Data, el)
	}

	return result, nil
}

func (s *SubsRepository) FindUserTopics(uid string) ([]string, error) {

	userId, err := primitive.ObjectIDFromHex(uid)
	if err != nil {
		return nil, err
	}

	project := bson.D{{"topic", 1}}
	filter := bson.D{{
		Key: "user_id", Value: userId,
	}}

	cur, err := s.coll.Find(context.TODO(), filter, &options.FindOptions{
		Projection: project,
	})
	if err != nil {

		if errors.As(err, &mongo.ErrNoDocuments) {
			return []string{}, nil
		}

		return nil, err
	}

	results := make([]string, 0)

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem models.UserSubModel
		err := cur.Decode(&elem)
		if err != nil {
			continue
		}

		results = append(results, elem.Topic)
	}

	return results, nil
}
