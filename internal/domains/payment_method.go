package models

type PaymentMethod struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
