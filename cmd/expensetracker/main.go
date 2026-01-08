package main

import "ExpenseTracker/internal/expenses"

func main() {
	el := expenses.NewExpenseList()

	el.ShowList()
}
