package cli

import (
	"fmt"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func AddExpense(category string, description string, amount float64) error {

	if description == "" || category == "" {
		return fmt.Errorf("Описание,категория не может быть пустым")
	}
	if amount < 0 {
		return fmt.Errorf("Расход не может быть меньше 0")
	}
	exp, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("Ошибка загрзуки файла %w", err)
	}
	newID := 1
	for _, e := range exp {
		if e.ID >= newID {
			newID = e.ID + 1
		}
	}
	newExp := model.NewExpense(newID, description, amount, category)
	exp = append(exp, newExp)
	if err := storage.SaveExpenses(exp); err != nil {
		return fmt.Errorf("Ошибка сохранения файла %w", err)
	}
	fmt.Printf("Расход успешно добавлен (ID: %d)\n", newID)
	return nil
}
