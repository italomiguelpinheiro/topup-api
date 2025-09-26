package repositories

import (
	"database/sql"
	"fmt"
	models "topup-api/internal/domains"
)

type PurchaseRepository struct {
	connection *sql.DB
}

func NewPurchaseRepository(connection *sql.DB) PurchaseRepository {
	return PurchaseRepository{
		connection: connection,
	}
}

func (r *PurchaseRepository) CreateOrder(order models.Order) (int, error) {

	var id int
	query, err := r.connection.Prepare("INSERT INTO orders" +
		"(id, order_number, client_id, offer_id, amount, payment_method_id, created_at)" +
		" VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(order.ID, order.OrderNumber, order.ClientID, order.OfferID, order.Amount, order.PaymentMethodID, order.CreatedAt).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil

}
