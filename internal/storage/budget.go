package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Budget struct {
	Month  int     `json:"month"`
	Amount float64 `json:"amount"`
}

const budgetFile = "data/budget.json"

func SaveBudget(b Budget) error {
	data, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации бюджета: %w", err)
	}
	return os.WriteFile(budgetFile, data, 0644)
}

func LoadBudget() (Budget, error) {
	var b Budget
	path := filepath.Join(".", budgetFile)
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return b, nil
	}
	if err != nil {
		return b, fmt.Errorf("ошибка чтения файла бюджета: %w", err)
	}
	if err := json.Unmarshal(data, &b); err != nil {
		return b, fmt.Errorf("ошибка разбора бюджета: %w", err)
	}
	return b, nil
}
