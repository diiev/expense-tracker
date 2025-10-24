package model

import "time"

type Expense struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:category`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:createdAt`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewExpense(id int, description string, amount float64, category string) *Expense {
	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Category:    category,
		Date:        time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

}
