package main

import (
	"ExpenseTracker/internal/expenses"
	"ExpenseTracker/internal/storage"
	"fmt"
)

const path = "storage/expenses.json"

func main() {
	el := expenses.NewExpenseList()
	el.AddExpense("Апельсин", 20.0)
	el.AddExpense("Банана", 39.2)

	storage.SaveExpenses(path, el)
	data, err := storage.LoadExpenses(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
