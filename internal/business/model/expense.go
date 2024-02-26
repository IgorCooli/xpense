package model

import "time"

type Expense struct {
	ID           string
	Value        float32
	PaymentDate  time.Time
	Installments uint
}
