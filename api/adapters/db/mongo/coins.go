package mongo

import (
	"context"
	"fmt"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"crypto-viewer/pkg/clients"
	"crypto-viewer/src/entities/dtos"
)

type storageMongo struct {
	collection *mongo.Collection
	logger     zerolog.Logger
}

func NewStorageMG(collection string, database *mongo.Database, l zerolog.Logger) clients.Storage {
	return &storageMongo{
		collection: database.Collection(collection),
		logger:     l,
	}
}

func (d storageMongo) Create(ctx context.Context, coin dtos.Coin) error {
	result, err := d.collection.InsertOne(ctx, coin)
	if err != nil {
		d.logger.Error().Err(err).Msg("Didn't created coin instance")

		return err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		d.logger.Debug().Msgf("coins with id (%v) saved", oid.Hex())

		return nil
	}

	d.logger.Error().Err(err).Msg("failed to convert object_id to hex")

	return err
}

func (d storageMongo) FindOne(ctx context.Context, id string) (coin dtos.Coin, err error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		d.logger.Error().Err(err).Msgf("failed to convert hex to objectid. hex: %s", id)

		return coin, err
	}

	filter := bson.M{"_id": oid}

	result := d.collection.FindOne(ctx, filter)

	if result.Err() != nil {
		d.logger.Error().Err(err).Msgf("failed to find one user by id: %s", id)

		return coin, err
	}

	if err = result.Decode(&coin); err != nil {
		d.logger.Error().Err(err).Msgf("failed to decode user by id: %s", id)

		return coin, err
	}

	return coin, nil
}

func (d storageMongo) Update(ctx context.Context, c dtos.Coin) error {
	filter := bson.M{"_id": c.ID}

	userBytes, err := bson.Marshal(c)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to marshal user")

		return err
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to unmarshal user bytes")

		return err
	}

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := d.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to execute update user query")

		return err
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}

func (d storageMongo) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to convert user ID to ObjectID")

		return err
	}

	filter := bson.M{"_id": objectID}

	result, err := d.collection.DeleteOne(ctx, filter)
	if err != nil {
		d.logger.Error().Err(err).Msg("failed to execute query")

		return err
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}

func (d storageMongo) FindAll(ctx context.Context) (coins []dtos.Coin, err error) {
	cursor, err := d.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		d.logger.Error().Err(err).Msg("failed to find all coins")

		return coins, err
	}

	if err = cursor.All(ctx, &coins); err != nil {
		d.logger.Error().Err(err).Msg("failed to read all documents from cursor")

		return coins, err
	}

	return coins, nil
}
