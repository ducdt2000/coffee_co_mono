package purchase

import (
	"context"

	"github.com/Rhymond/go-money"
)

type CardChargeService interface {
	ChargeCard(ctx context.Context, amount money.Money, cardToken string) error
}
