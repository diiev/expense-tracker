package cli

import (
	"encoding/csv"
	"fmt"
	"os"

	"go.mod/internal/storage"
)

func ExportToCSV(filename, category string) error {
	expenses, err := storage.LoadExpense()
	if err != nil {
		return fmt.Errorf("ошибка загрузки данных: %w", err)
	}

	file, err := os.Create("data/" + filename)
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"ID", "Дата", "Категория", "Описание", "Сумма"}
	if err := writer.Write(headers); err != nil {
		return fmt.Errorf("ошибка записи заголовков: %w", err)
	}

	for _, e := range expenses {

		if category != "" && e.Category != category {
			continue
		}

		record := []string{
			fmt.Sprintf("%d", e.ID),
			e.Date.Format("2006-01-02"),
			e.Category,
			e.Description,
			fmt.Sprintf("%.2f", e.Amount),
		}

		if err := writer.Write(record); err != nil {
			return fmt.Errorf("ошибка записи строки: %w", err)
		}
	}

	fmt.Printf("✅ Экспорт завершён: %s\n", filename)
	return nil
}
