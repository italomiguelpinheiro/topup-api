package usecases

import (
	models "topup-api/internal/domains"
	"topup-api/internal/repositories"
)

type OfferUsecase struct {
	repository repositories.OfferRepository
}

func NewOfferUsecase(repo repositories.OfferRepository) OfferUsecase {
	return OfferUsecase{
		repository: repo,
	}
}

func (ou *OfferUsecase) GetOffers() ([]models.Offer, error) {
	return ou.repository.GetOffers()
}
