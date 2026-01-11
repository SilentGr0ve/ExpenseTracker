package storage

import (
	"ExpenseTracker/internal/expenses"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveExpenses(path string, el expenses.ExpenseList) error {
	dir := filepath.Dir(path)
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}

	jsonData, err := json.MarshalIndent(el, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadExpenses(path string) (expenses.ExpenseList, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return expenses.NewExpenseList(), nil
	}

	jsonData, err := os.ReadFile(path)
	if err != nil {
		return expenses.ExpenseList{}, err
	}

	if len(jsonData) == 0 {
		return expenses.ExpenseList{}, fmt.Errorf("expenses file is empty")
	}

	expensesList := expenses.ExpenseList{}

	err = json.Unmarshal(jsonData, &expensesList)
	if err != nil {
		return expenses.ExpenseList{}, err
	}
	return expensesList, nil
}
