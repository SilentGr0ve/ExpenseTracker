package cli

import (
	"ExpenseTracker/internal/expenses"
	"flag"
)

func RunHelp() {
	PrintHelp()
}

func RunAdd(args []string, el *expenses.ExpenseList) error {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	desc := fs.String("description", "", "expense description")
	amount := fs.Float64("amount", 0, "amount of expense")
	fs.Parse(args)

	err := el.AddExpense(*desc, *amount)
	if err != nil {
		return err
	}
	return nil
}
