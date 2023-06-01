package purchase

import (
	coffee_co "coffee_co/internal"
	"coffee_co/internal/loyalty"
	"coffee_co/internal/payment"
	"coffee_co/internal/store"
	"context"
	"errors"
	"fmt"
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
	storeService      StoreService
	purchaseRepo      Repository
}

func NewService(cardChargeService CardChargeService, storeService StoreService, purchaseRepo Repository) *Service {
	return &Service{cardChargeService: cardChargeService, storeService: storeService, purchaseRepo: purchaseRepo}
}

func (s Service) CompletePurchase(ctx context.Context, purchase *Purchase, coffeeBuxCard *loyalty.CoffeeBux) error {
	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}

	if err := s.calculateStoreSpecificDiscount(ctx, purchase.Store.ID, purchase); err != nil {
		return err
	}

	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		if err := s.cardChargeService.ChargeCard(ctx, purchase.Total, purchase.CardToken); err != nil {
			return errors.New("card charge failed, cancelling purchase")
		}
	case payment.MEANS_CASH:
		//TODO: for reader to add
	case payment.MEANS_COFFEBUX:
		if err := coffeeBuxCard.Pay(ctx, purchase.ProductsToPurchase); err != nil {
			return fmt.Errorf("failed to charge loyalty card: %w", err)
		}
	default:
		return errors.New("unknown payment type")
	}
	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return errors.New("failed to store purchase")
	}

	if coffeeBuxCard != nil && purchase.PaymentMeans != payment.MEANS_COFFEBUX {
		coffeeBuxCard.AddStamp()
	}

	return nil
}

func (s Service) calculateStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID, purchase *Purchase) error {
	discount, err := s.storeService.GetStoreSpecificDiscount(ctx, storeID)
	if err != nil && err != store.ErrNoDiscount {
		return fmt.Errorf("failed to get discount: %w", err)
	}
	purchasePrice := purchase.Total
	if discount > 0 {
		purchase.Total = *purchasePrice.Multiply(int64(100 - discount))
	}
	return nil
}
