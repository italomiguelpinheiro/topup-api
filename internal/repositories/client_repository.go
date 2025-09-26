package repositories

import (
	"database/sql"
	"fmt"
	models "topup-api/internal/domains"
)

type ClientRepository struct {
	connection *sql.DB
}

func NewClientRepository(connecyion *sql.DB) ClientRepository {
	return ClientRepository{
		connection: connecyion,
	}
}

func (r *ClientRepository) GetClients() ([]models.Client, error) {
	query := "SELECT id, msisdn, balance, credit_limit FROM clients"
	rows, err := r.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []models.Client{}, err
	}

	var clientList []models.Client
	var clientObj models.Client

	for rows.Next() {
		err = rows.Scan(
			&clientObj.ID,
			&clientObj.MSISDN,
			&clientObj.Balance,
			&clientObj.CreditLimit)

		if err != nil {
			fmt.Println(err)
			return []models.Client{}, err
		}

		clientList = append(clientList, clientObj)
	}

	rows.Close()
	return clientList, nil
}
