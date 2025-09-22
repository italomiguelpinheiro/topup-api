package models

type CreditCard struct {
	ID     int64  `json:"id"`
	Number string `json:"card_number"`
	Brand  string `json:"brand"`
}
