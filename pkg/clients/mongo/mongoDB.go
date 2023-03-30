package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"crypto-viewer/pkg/clients"
	"crypto-viewer/pkg/logger"
	"crypto-viewer/src/entities"
)

type db struct {
	collection *mongo.Collection
}

func (d db) Create(ctx context.Context, coins entities.Coin) (string, error) {
	result, err := d.collection.InsertOne(ctx, coins)
	if err != nil {
		logger.Log.Error().Msg("Didnt created")
		return "", err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return oid.Hex(), nil
	}

	return "", err
}

func (d db) Read(ctx context.Context, id string) (c entities.Coin, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c, err
	}

	filter := bson.M{"id": oid}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		return c, fmt.Errorf("ddd")
	}

	if err = result.Decode(&c); err != nil {
		return c, fmt.Errorf("ddd")
	}

	return c, nil
}

func (d db) Update(ctx context.Context, coins entities.Coin) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (d db) Delete(ctx context.Context, coins entities.Coin) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewStorage(collection string, database *mongo.Database) clients.Storage {
	database.Collection(collection)
	return &db{
		collection: database.Collection(collection),
	}
}
