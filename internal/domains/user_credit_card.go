package models

type UserCreditCard struct {
	ID       int64 `json:"id"`
	ClientID int64 `json:"client_id"`
	CardID   int64 `json:"card_id"`
}
