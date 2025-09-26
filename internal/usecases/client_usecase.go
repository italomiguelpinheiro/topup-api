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

func (u *ClientUsecase) GetClientById(client_id int) (*models.Client, error) {
	client, err := u.repository.GetClientById(client_id)

	if err != nil {
		return nil, err
	}

	return client, nil
}
