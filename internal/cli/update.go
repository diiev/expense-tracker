package cli

import (
	"fmt"

	"go.mod/internal/storage"
)

func UpdateExp(id int, category string, description string, amount float64) error {
	expenses, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("ошибка загрузки данных: %w", err)
	}

	found := false

	for _, e := range expenses {
		if e.ID == id {
			found = true

			if category != "" {
				e.Category = category
			}
			if description != "" {
				e.Description = description
			}
			if amount > 0 {
				e.Amount = amount
			}

			break
		}
	}

	if !found {
		fmt.Printf("Запись с ID-%d не найдена\n", id)
		return nil
	}

	if err := storage.SaveExpenses(expenses); err != nil {
		return fmt.Errorf("ошибка сохранения данных: %w", err)
	}

	fmt.Printf("Запись с ID-%d успешно обновлена\n", id)
	return nil
}
