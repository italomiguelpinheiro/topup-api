package usecases

import (
	models "topup-api/internal/domains"
	"topup-api/internal/repositories"
)

type PurchaseUsecase struct {
	repository repositories.PurchaseRepository
}

func NewPurchaseUsecase(repo repositories.PurchaseRepository) PurchaseUsecase {
	return PurchaseUsecase{
		repository: repo,
	}
}

func (pu *PurchaseUsecase) CreateOrder(order models.Order) (models.Order, error) {
	orderId, err := pu.repository.CreateOrder(order)

	if err != nil {
		return models.Order{}, err
	}

	order.ID = int64(orderId)
	return order, nil
}
