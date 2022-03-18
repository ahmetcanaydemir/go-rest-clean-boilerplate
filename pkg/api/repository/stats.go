package repository

import (
	"context"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/utils/helper"
	"go.uber.org/zap"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/db"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	statsCollectionName = "Stats"
)

type StatsRepository interface {
	GetLastStats() []model.Stats
}

type statsRepository struct {
	collection *mongo.Collection
	gr         GenericRepository
}

func NewStatsRepository() StatsRepository {
	collection := db.GetDatabase().DB.Collection(statsCollectionName)
	return &statsRepository{
		collection: collection,
		gr:         NewGenericRepository(collection),
	}
}

func (r statsRepository) GetLastStats() []model.Stats {
	var stats []model.Stats

	opts := options.Find().SetSort(bson.M{"_id": -1}).SetLimit(1)
	cursor := r.gr.Find(bson.D{}, opts)

	for cursor.Next(context.TODO()) {
		var stat model.Stats

		err := bson.UnmarshalWithRegistry(helper.CustomRegistry, cursor.Current, &stat)
		if err != nil {
			zap.L().Error("Error while decoding", zap.Error(err))
			panic(err)
		}
		stats = append(stats, stat)
	}
	return stats
}
