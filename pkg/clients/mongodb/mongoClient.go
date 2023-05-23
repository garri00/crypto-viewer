package mongodb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"crypto-viewer/pkg/logger"
	"crypto-viewer/src/config"
)

func NewClient(ctx context.Context, configs config.MongoDBConf) (db *mongo.Database, err error) {
	var mongoDBURL string
	mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", configs.Username, configs.Password, configs.Host, configs.Port)

	opts := options.Client().ApplyURI(mongoDBURL)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		logger.Log.Err(err).Msg("failed to connect mongoDB")
		return nil, err
	}

	var result bson.M
	if err := client.Database(configs.Database).RunCommand(ctx, bson.D{{"ping", 1}}).Decode(&result); err != nil {
		logger.Log.Err(err).Msg("failed to ping mongoDB")
		return nil, err
	}

	logger.Log.Info().Msgf("successfully connected to MongoDB host: %s, database: %s", configs.Host, configs.Database)

	return client.Database(configs.Database), nil
}
