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
			fmt.Println("-------------------------------")
			fmt.Printf("Номер задачи: %d\nОписание: %s\nСтатус: %s\nСоздана: %s\nОбновлено: %s\n",
				t.ID, t.Description, t.Category, t.Date.Format("2006-01-02 15:04"), t.Amount)
		}
	}
	if len(filtered) == 0 {
		fmt.Println("Задачи не найдены")
	}
	return nil
}
