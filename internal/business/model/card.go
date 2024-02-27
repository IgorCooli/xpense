package model

type Card struct {
	ID         string `json:"id"`
	CardNumber uint   `json:"cardNumber"`
	CardBrand  string `json:"cardBrand"`
	UserID     string `json:"userId"`
}
