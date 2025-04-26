package models

type Transaction struct {
	AccountID int     `json:"account_id"`
	Type      string  `json:"tyoe"`
	Amount    float64 `json:"amount"`
	TargetID  int     `json:"target_id"`
}