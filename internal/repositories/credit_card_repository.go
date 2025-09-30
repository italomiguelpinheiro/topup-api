package repositories

import (
	"database/sql"
	"fmt"
	models "topup-api/internal/domains"
)

type CreditCardRepository struct {
	connection *sql.DB
}

func NewCreditCardRepository(connection *sql.DB) CreditCardRepository {
	return CreditCardRepository{
		connection: connection,
	}
}

func (r *CreditCardRepository) GetCreditCardsByClientId(clientId int) ([]models.CreditCard, error) {
	query := `
		SELECT cc.id, cc.card_number, cc.brand
		FROM user_credit_cards ucc
		INNER JOIN credit_cards cc ON cc.id = ucc.card_id
		WHERE ucc.client_id = $1
	`

	rows, err := r.connection.Query(query, clientId)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer rows.Close()

	var cards []models.CreditCard

	for rows.Next() {
		var card models.CreditCard
		err := rows.Scan(&card.ID, &card.Number, &card.Brand)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		cards = append(cards, card)
	}

	return cards, nil

}
