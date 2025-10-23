package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"go.mod/internal/model"
)

const FileName = "data/expense.json"

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
