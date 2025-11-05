package cli

import (
	"fmt"
	"time"

	"go.mod/internal/storage"
)

func getMonthString(month int) string {
	mapMonth := map[int]string{
		1:  "Январь",
		2:  "Февраль",
		3:  "Март",
		4:  "Апрель",
		5:  "Май",
		6:  "Июнь",
		7:  "Июль",
		8:  "Август",
		9:  "Сентябрь",
		10: "Октябрь",
		11: "Ноябрь",
		12: "Декабрь",
	}
	return mapMonth[month]
}
func SummaryExp(month int) error {
	exp, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("Ошибка загрузки файла %w", err)
	}
	var summary float64
	for _, e := range exp {
		if e.Date.Month() == time.Month(month) && e.Date.Year() == time.Now().Year() || month == 0 {
			summary += e.Amount
		}

	}
	if month == 0 {
		fmt.Printf("Общий расход за все время равен %.2f руб.\n", summary)
	} else {

		fmt.Printf("Общий расход за %s месяц равен %.2f руб.\n", getMonthString(month), summary)
	}

	budget, _ := storage.LoadBudget()
	if budget.Month == month && budget.Amount > 0 {
		fmt.Printf("💰 Бюджет: %.2f руб.\n", budget.Amount)
		if summary > budget.Amount {
			fmt.Printf("⚠️  Превышение бюджета на %.2f руб.\n", summary-budget.Amount)
		} else {
			fmt.Printf("Остаток бюджета: %.2f руб.\n", budget.Amount-summary)
		}
	}

	return nil
}
