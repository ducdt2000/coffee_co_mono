package payment

import (
	"context"
	"errors"

	"github.com/Rhymond/go-money"
	"github.com/stripe/stripe-go/v73/client"
)

type StripeService struct {
	stripeClient *client.API
}

func NewStripeService(apiKey string) (*StripeService, error) {
	if apiKey == "" {
		return nil, errors.New("API key cannot be nil")
	}

	sc := &client.API{}

	sc.Init(apiKey, nil)
	return &StripeService{stripeClient: sc}, nil
}

func (s StripeService) ChargeCard(ctx context.Context, amount money.Money, cardToken string) error {
	//here
}
