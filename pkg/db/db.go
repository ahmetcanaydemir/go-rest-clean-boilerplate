package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/configs"
	"go.uber.org/zap"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	dbName           = "mydb"
	mongoConnTimeout = 10 * time.Second
)

type Database struct {
	DB     *mongo.Database
	Client *mongo.Client
}

var database *Database
var once sync.Once

func GetDatabase() *Database {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), mongoConnTimeout)
		defer cancel()

		client, err := mongo.Connect(ctx, options.Client().ApplyURI(configs.GetConfigs().DbConnectionString))
		if err != nil {
			zap.L().Fatal("Failed to connect to mongo", zap.Error(err))
			panic(fmt.Sprintf("failed to connect to mongo: %s", err))
		}

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			zap.L().Fatal("Could not ping mongodb after connect", zap.Error(err))
			panic(fmt.Sprintf("could not ping mongodb after connect: %s", err))
		}
		mongoDB := client.Database(dbName)

		database = &Database{
			DB:     mongoDB,
			Client: client,
		}
	})

	return database
}

func (r *Database) Close() error {
	return r.Client.Disconnect(context.TODO())
}
