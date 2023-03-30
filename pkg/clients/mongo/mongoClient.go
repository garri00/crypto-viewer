package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crypto-viewer/pkg/logger"
	"crypto-viewer/src/config"
)

func NewClient(ctx context.Context, configs config.MongoDBConf) (db *mongo.Database, err error) {
	var mongoDBURL string
	mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", configs.Username, configs.Password, configs.Host, configs.Port)

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Log.Error().Msg("failed to connect to mongoDB")
		return nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		logger.Log.Error().Msg("failed to ping to mongoDB")
		return nil, err
	}

	return client.Database(configs.Database), nil
}
