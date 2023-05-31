package store

import (
	coffee_co "coffee_co/internal"

	"github.com/google/uuid"
)

type Store struct {
	ID              uuid.UUID
	Location        string
	ProductsForSale []coffee_co.Product
}
