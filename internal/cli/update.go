package cli

import (
	"fmt"

	"go.mod/internal/storage"
)

func UpdateExp(id int, category string, description string, amount float64) error {
	exp, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("Ошибка загрузки данных")
	}
	found := false

	// var newExp []*model.Expense
	for _, e := range exp {
		if e.ID == id {
			found = true
			if category != "" && description != "" && amount > 0 {
				e.Category = category
				e.Description = description
				e.Amount = amount
			}
			if category == "" && description == "" {
				e.Amount = amount
			}
		}
	}
	if err := storage.SaveExpenses(exp); err != nil {
		return fmt.Errorf("Ошибка сохранения данных %w", err)

	}
	if found {
		fmt.Printf("Запись с ID-%d есть, выберите что хотите поменять?\n1-Категорию\n2-Описание\n3-Сумму\n", id)
	} else {
		fmt.Printf("Запись с ID-%d не найдена\n", id)
	}
	return nil
}
