package main

import "ExpenseTracker/expenses"

func main() {
	el := expenses.NewExpenseList()
	_ = el.AddExpense("Апельсин", 20)
	_ = el.AddExpense("Банан", 10)
	_ = el.AddExpense("Яблоко", 5)
	_ = el.RemoveExpense()
	_ = el.ShowList()

}
