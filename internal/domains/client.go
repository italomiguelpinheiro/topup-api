package models

type Client struct {
	ID          int64   `json:"id"`
	MSISDN      string  `json:"msisdn"`
	Balance     float64 `json:"balance"`
	CreditLimit float64 `json:"credit_limit"`
}
