package usecases

import (
	models "topup-api/internal/domains"
	"topup-api/internal/repositories"
)

type CreditCardUsecase struct {
	repository repositories.CreditCardRepository
}

func NewCreditCardUsecase(repo repositories.CreditCardRepository) CreditCardUsecase {
	return CreditCardUsecase{
		repository: repo,
	}
}

func (u *CreditCardUsecase) GetCreditCardsByClientId(clientID int) ([]models.CreditCard, error) {
	return u.repository.GetCreditCardsByClientId(clientID)
}
