package usecases

import (
	models "topup-api/internal/domains"
	"topup-api/internal/repositories"
)

type ClientUsecase struct {
	repository repositories.ClientRepository
}

func NewClientUsercase(repo repositories.ClientRepository) ClientUsecase {
	return ClientUsecase{
		repository: repo,
	}
}

func (u *ClientUsecase) GetClients() ([]models.Client, error) {
	return u.repository.GetClients()
}
