package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"go.mod/internal/model"
)

const FileName = "data/expense.json"

// Загружает расходы из файла, если нет файла создает его
func LoadExpense() ([]*model.Expense, error) {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		empty := []*model.Expense{}
		data, err := json.MarshalIndent(empty, "", " ")
		if err != nil {
			return nil, fmt.Errorf("Ошибка сериализации %w", err)
		}
		if err := os.WriteFile(FileName, data, 0644); err != nil {
			return nil, fmt.Errorf("Ошибка создания файла %w", err)
		}
		return empty, nil
	}
	data, err := os.ReadFile(FileName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка чтения файла: %w", err)
	}
	if len(data) == 0 {
		return []*model.Expense{}, nil
	}
	var expenses []*model.Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, fmt.Errorf("Ошибка десереиализации файла %w", err)
	}
	return expenses, nil
}

// Cохраняет изменения в файл

func SaveExpenses(expenses []*model.Expense) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return fmt.Errorf("Ошибка записи данных %w", err)
	}
	tmpFile := FileName + ".tmp"
	f, err := os.Create(tmpFile)
	if err != nil {
		return fmt.Errorf("Ошибка создания файла %w", err)
	}
	_, err = f.Write(data)
	if err != nil {
		f.Close()
		return fmt.Errorf("Ошибка записи данных в файл %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("Ошбика закрытия файла %w", err)
	}
	if err := os.Rename(tmpFile, FileName); err != nil {
		return fmt.Errorf("Ошибка переименования файла")
	}
	defer func() {
		f.Close()
	}()
	return nil
}
