package loyalty

import (
	coffee_co "coffee_co/internal"
	"coffee_co/internal/store"

	"github.com/google/uuid"
)

type CoffeeBux struct {
	ID                                  uuid.UUID
	Store                               store.Store
	CoffeeLover                         coffee_co.CoffeeLover
	FreeDrinksAvailable                 int
	RemainingDrinkPurchaseUtilFreeDrink int
}
