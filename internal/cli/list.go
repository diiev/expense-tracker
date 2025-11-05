package cli

import (
	"fmt"
	"strings"

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
		if category == e.Category {
			filtered = append(filtered, e)
		}

	}
	fmt.Printf("%-4s | %-12s | %-15s | %-20s | %-15s |\n", "ID", "Дата", "Категория", "Описание", "Сумма")
	fmt.Println(strings.Repeat("-", 75))

	for _, t := range filtered {
		if t != nil {
			fmt.Printf("%-4d | %-12s | %-15s | %-20s | %-10.2f руб. |\n",
				t.ID,
				t.Date.Format("2006-01-02"),
				t.Category,
				t.Description,
				t.Amount)
		}
	}

	if len(filtered) == 0 {
		fmt.Println("Задачи не найдены")
	}

	return nil
}
