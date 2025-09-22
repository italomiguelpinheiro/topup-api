package models

import "time"

type Order struct {
	ID              int64     `json:"id"`
	OrderNumber     string    `json:"order_number"`
	ClientID        int64     `json:"client_id"`
	OfferID         *int64    `json:"offer_id"` // nullable
	Amount          float64   `json:"amount"`
	PaymentMethodID int64     `json:"payment_method_id"`
	CreatedAt       time.Time `json:"created_at"`
}
