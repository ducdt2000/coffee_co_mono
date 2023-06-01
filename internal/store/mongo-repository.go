package store

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	discountCollection *mongo.Collection
}

func NewMongoRepository(ctx context.Context, connStr string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}
	discountCollection := client.Database("coffee_co").Collection("store_discounts")
	return &MongoRepository{
		discountCollection: discountCollection,
	}, nil
}

func (mr MongoRepository) GetStoreDiscount(ctx context.Context, storeID uuid.UUID) (float32, error) {
	var discount float32

	if err := mr.discountCollection.FindOne(
		ctx, bson.D{{Key: "store_id", Value: storeID.String()}},
	).Decode(&discount); err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, ErrNoDiscount
		}
		return 0, fmt.Errorf("failed to find discount for store: %w", err)
	}

	return discount, nil
}
