package cli

import (
	"fmt"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func DeleteExpense(id int) error {
	exp, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("Ошибка загрузки файла %w", err)
	}
	found := false
	var newExp []*model.Expense
	for _, e := range exp {
		if e.ID != id {
			found = true
			newExp = append(newExp, e)
		}

	}
	if err := storage.SaveExpenses(newExp); err != nil {
		return fmt.Errorf("Ошибка сохранения файла %w", err)
	}
	if found {
		fmt.Printf("Расход с id %d удален\n", id)
	} else {
		fmt.Printf("Расход с id %d не найден\n", id)
	}
	return nil
}
