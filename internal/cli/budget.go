package cli

import (
	"fmt"

	"go.mod/internal/storage"
)

func SetBudget(month int, amount float64) error {
	if month < 1 || month > 12 {
		return fmt.Errorf("Месяц должен быть в диапазоне 1–12")
	}
	if amount <= 0 {
		return fmt.Errorf("Сумма бюджета должна быть больше нуля")
	}

	b := storage.Budget{
		Month:  month,
		Amount: amount,
	}

	if err := storage.SaveBudget(b); err != nil {
		return err
	}

	fmt.Printf("✅ Установлен бюджет на %d месяц: %.2f руб.\n", month, amount)
	return nil
}
