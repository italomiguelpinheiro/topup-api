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

func (r *ClientRepository) GetClientById(client_id int) (*models.Client, error) {
	query, err := r.connection.Prepare("SELECT * FROM clients WHERE id = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var client models.Client

	err = query.QueryRow(client_id).Scan(
		&client.ID,
		&client.MSISDN,
		&client.Balance,
		&client.CreditLimit,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()
	return &client, nil
}
