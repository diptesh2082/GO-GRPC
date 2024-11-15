package models

type Billing struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}