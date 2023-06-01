package store

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

var ErrNoDiscount = errors.New("no discount for store")

type Repository interface {
	GetStoreDiscount(ctx context.Context, storeID uuid.UUID) (float32, error)
}
