package cli

import (
	"fmt"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func ShowExpense() error {
	expenses, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("Ошибка загзрузки файла %w", err)
	}
	var filtered []*model.Expense

}
