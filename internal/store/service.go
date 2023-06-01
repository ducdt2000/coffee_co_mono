package store

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s Service) GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (float32, error) {
	discount, err := s.repository.GetStoreDiscount(ctx, storeID)
	if err != nil {
		return 0, err
	}
	return discount, nil
}
