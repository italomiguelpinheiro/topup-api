package repositories

import (
	"database/sql"
	"fmt"
	models "topup-api/internal/domains"
)

type OfferRepository struct {
	connection *sql.DB
}

func NewOfferRepository(connection *sql.DB) OfferRepository {
	return OfferRepository{
		connection: connection,
	}
}

func (r *OfferRepository) GetOffers() ([]models.Offer, error) {
	query := "SELECT id, name, value FROM offers"
	rows, err := r.connection.Query(query)

	if err != nil {
		fmt.Println(err)
		return []models.Offer{}, err
	}

	var offerList []models.Offer
	var offerObj models.Offer

	for rows.Next() {
		err = rows.Scan(
			&offerObj.ID,
			&offerObj.Name,
			&offerObj.Value)

		if err != nil {
			fmt.Println(err)
			return []models.Offer{}, err
		}

		offerList = append(offerList, offerObj)
	}

	rows.Close()
	return offerList, nil

}
