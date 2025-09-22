package models

type Offer struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}
