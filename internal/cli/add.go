package cli

import (
	"fmt"
	"time"

	"go.mod/internal/model"
	"go.mod/internal/storage"
)

func AddExpense(category string, description string, amount float64) error {
	if description == "" || category == "" {
		return fmt.Errorf("описание и категория не могут быть пустыми")
	}
	if amount <= 0 {
		return fmt.Errorf("расход должен быть больше нуля")
	}

	expenses, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("ошибка загрузки файла: %w", err)
	}

	newID := 1
	for _, e := range expenses {
		if e.ID >= newID {
			newID = e.ID + 1
		}
	}

	newExp := model.NewExpense(newID, description, amount, category)
	expenses = append(expenses, newExp)
	if err := storage.SaveExpenses(expenses); err != nil {
		return fmt.Errorf("ошибка сохранения файла: %w", err)
	}

	fmt.Printf("✅ Расход успешно добавлен (ID: %d)\n", newID)

	currentMonth := int(time.Now().Month())
	budget, _ := storage.LoadBudget()

	if budget.Month == currentMonth && budget.Amount > 0 {
		total := 0.0
		for _, e := range expenses {
			if int(e.Date.Month()) == currentMonth {
				total += e.Amount
			}
		}

		if total > budget.Amount {
			fmt.Printf("⚠️  Внимание: вы превысили бюджет на %.2f руб.\n", total-budget.Amount)
		} else {
			fmt.Printf("💰 Остаток бюджета: %.2f руб.\n", budget.Amount-total)
		}
	}

	return nil
}
