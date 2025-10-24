package cli

import (
	"fmt"

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
		summary += e.Amount
	}
	fmt.Printf("Общий расход за %s месяц равен %.2f руб.\n", getMonthString(month), summary)
	return nil
}
