package main

import (
	coffee_co "coffee_co/internal"
	"coffee_co/internal/payment"
	"coffee_co/internal/purchase"
	"coffee_co/internal/store"
	"context"
	"log"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

func main() {
	ctx := context.Background()

	stripeTestAPIKey := "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	cardToken := "tok_visa_debit"

	mongoConString := "mongodb://root:example@localhost:27017"

	stripeService, err := payment.NewStripeService(stripeTestAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	purchaseRepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}

	storeRepo, err := store.NewMongoRepository(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}

	storeService := store.NewService(storeRepo)

	purchaseService := purchase.NewService(stripeService, storeService, purchaseRepo)

	someStoreId := uuid.New()

	purchase := &purchase.Purchase{
		CardToken: cardToken,
		Store: store.Store{
			ID: someStoreId,
		},
		ProductsToPurchase: []coffee_co.Product{
			{ItemName: "item1", BasePrice: *money.New(3300, "USD")},
		},
		PaymentMeans: payment.MEANS_CARD,
	}

	if err := purchaseService.CompletePurchase(ctx, purchase, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("purchase was successful")
}
