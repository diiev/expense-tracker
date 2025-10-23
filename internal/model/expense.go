package model

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewExpense(id int, description string, amount float64) *Expense {
	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
		UpdatedAt:   time.Now(),
	}

}
