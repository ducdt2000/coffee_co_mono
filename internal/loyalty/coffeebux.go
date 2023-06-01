package loyalty

import (
	coffee_co "coffee_co/internal"
	"coffee_co/internal/store"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CoffeeBux struct {
	ID                                  uuid.UUID
	Store                               store.Store
	CoffeeLover                         coffee_co.CoffeeLover
	FreeDrinksAvailable                 int
	RemainingDrinkPurchaseUtilFreeDrink int
}

func (cb *CoffeeBux) AddStamp() {
	if cb.RemainingDrinkPurchaseUtilFreeDrink == 1 {
		cb.RemainingDrinkPurchaseUtilFreeDrink = 10
		cb.FreeDrinksAvailable += 1
	} else {
		cb.RemainingDrinkPurchaseUtilFreeDrink--
	}
}

func (cb *CoffeeBux) Pay(ctx context.Context, products []coffee_co.Product) error {
	lenOfProduct := len(products)
	if lenOfProduct == 0 {
		return errors.New("nothing to buy")
	}
	if cb.FreeDrinksAvailable < lenOfProduct {
		return fmt.Errorf("not enough coffeeBux to cover entire purchase. Have %d, need%d", len(products), cb.FreeDrinksAvailable)
	}

	cb.FreeDrinksAvailable -= lenOfProduct
	return nil
}
