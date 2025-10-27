package cli

import (
	"fmt"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func ShowExpense(category string) error {
	expenses, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("Ошибка загзрузки файла %w", err)
	}
	var filtered []*model.Expense
	for _, e := range expenses {
		if category == "" {
			filtered = append(filtered, e)
		}

	}
	for _, t := range filtered {
		if t != nil {
			fmt.Printf("%d  |\t%s |\t%s\t  |\t%s\t\t|\t%.2f руб.\t|\n", t.ID, t.Date.Format("2006-02-01"), t.Category, t.Description, t.Amount)
		}
	}
	if len(filtered) == 0 {
		fmt.Println("Задачи не найдены")
	}

	return nil
}
