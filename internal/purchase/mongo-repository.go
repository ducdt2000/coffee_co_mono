package purchase

import (
	coffee_co "coffee_co/internal"
	"coffee_co/internal/payment"
	"coffee_co/internal/store"
	"context"
	"fmt"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	purchaseCollection *mongo.Collection
}

func NewMongoRepo(ctx context.Context, connStr string) (*MongoRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))

	if err != nil {
		return nil, fmt.Errorf("failed to create a mongo client: %w", err)
	}

	purcharseCollection := client.Database("coffee_co").Collection("purchase")

	return &MongoRepository{purchaseCollection: purcharseCollection}, nil
}
func (mp *MongoRepository) Store(ctx context.Context, purchase Purchase) error {
	mongoPurchase := toMongoPurchase(purchase)
	_, err := mp.purchaseCollection.InsertOne(ctx, mongoPurchase)
	if err != nil {
		return fmt.Errorf("failed to persist purchase: %w", err)
	}
	return nil
}

type mongoPurchase struct {
	ID                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffee_co.Product
	Total              money.Money
	PaymentMeans       payment.Means
	TimeOfPurchase     time.Time
	CardToken          string
}

func toMongoPurchase(p Purchase) mongoPurchase {
	return mongoPurchase{
		ID:                 p.ID,
		Store:              p.Store,
		ProductsToPurchase: p.ProductsToPurchase,
		Total:              p.Total,
		PaymentMeans:       p.PaymentMeans,
		TimeOfPurchase:     p.TimeOfPurchase,
		CardToken:          p.CardToken,
	}
}
