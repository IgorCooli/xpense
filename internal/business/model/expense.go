package model

import "time"

type Expense struct {
	// ID           string    `json:"_id"`
	Value           float32         `json:"value"`
	PaymentDate     time.Time       `json:"paymentDate"`
	Installments    uint            `json:"installments"`
	Description     string          `json:"description"`
	Type            string          `json:"type"`
	Method          string          `json:"method"`
	ExpenseUserInfo ExpenseUserInfo `json:"expenseUserInfo"`
}
