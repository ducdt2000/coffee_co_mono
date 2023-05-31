package purchase

import (
	coffee_co "coffee_co/internal"
	"coffee_co/internal/payment"
	"coffee_co/internal/store"
	"context"
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Purchase struct {
	ID                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffee_co.Product
	Total              money.Money
	PaymentMeans       payment.Means
	TimeOfPurchase     time.Time
	CardToken          string
}

func (p *Purchase) validateAndEnrich() error {
	if len(p.ProductsToPurchase) == 0 {
		return errors.New("purchase must consist of at least one product")
	}
	p.Total = *money.New(0, "USD")
	for _, v := range p.ProductsToPurchase {
		newTotal, _ := p.Total.Add(&v.BasePrice)
		p.Total = *newTotal
	}

	if p.Total.IsZero() {
		return errors.New("likely mistake; purchase should never be 0. Please validate")
	}
	p.ID = uuid.New()
	p.TimeOfPurchase = time.Now()
	return nil
}

type Service struct {
	cardChargeService CardChargeService
	purchaseRepo      Repository
}

func (s Service) CompletePurchase(ctx context.Context, purchase *Purchase) error {
	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}

	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		if err := s.cardChargeService.ChargeCard(ctx, purchase.Total, *&purchase.CardToken); err != nil {
			return errors.New("card charge failed, cancelling purchase")
		}
	case payment.MEANS_CASH:
		//TODO: for reader to add
	default:
		return errors.New("unknown payment type")
	}
	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return errors.New("failed to store purchase")
	}

	return nil
}
