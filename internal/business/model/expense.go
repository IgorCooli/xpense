package model

import "time"

type Expense struct {
	ID           string    `json:"id"`
	Value        float32   `json:"value"`
	PaymentDate  time.Time `json:"paymentDate"`
	Month        string    `json:"month"`
	Year         string    `json:"year"`
	Installments uint      `json:"installments"`
	Description  string    `json:"description"`
	Type         string    `json:"type"`
	Method       string    `json:"method"`
	Card         Card      `json:"card"`
}
