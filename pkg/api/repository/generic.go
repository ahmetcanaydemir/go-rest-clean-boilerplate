package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type GenericRepository interface {
	FindOne(filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult
	Find(filter interface{}, opts *options.FindOptions) *mongo.Cursor
}

type genericRepository struct {
	collection *mongo.Collection
}

func NewGenericRepository(collection *mongo.Collection) GenericRepository {
	return &genericRepository{
		collection: collection,
	}
}

func (r genericRepository) FindOne(filter interface{}, opts *options.FindOneOptions) *mongo.SingleResult {
	var singleResult *mongo.SingleResult

	if opts == nil {
		singleResult = r.collection.FindOne(context.Background(), filter)
	} else {
		singleResult = r.collection.FindOne(context.Background(), filter, opts)
	}
	err := singleResult.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		zap.L().Error("Error while DB FindOne", zap.Error(err))
		panic(err)
	}

	return singleResult
}

func (r genericRepository) Find(filter interface{}, opts *options.FindOptions) *mongo.Cursor {
	var err error
	var cursor *mongo.Cursor

	if opts == nil {
		cursor, err = r.collection.Find(context.Background(), filter)
	} else {
		cursor, err = r.collection.Find(context.Background(), filter, opts)
	}

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		zap.L().Error("Error while mongodb find", zap.Error(err))
		panic(err)
	}
	return cursor
}
