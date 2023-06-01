package purchase

import (
	"context"

	"github.com/google/uuid"
)

type StoreService interface {
	GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (float32, error)
}
