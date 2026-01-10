package main

import (
	"ExpenseTracker/internal/cli"
	"ExpenseTracker/internal/expenses"
	"fmt"
)

func main() {
	el := expenses.NewExpenseList()
	err := cli.RunSummary(el)
	if err != nil {
		fmt.Println(err)
	}
	args := []string{
		"--description", "Апельсины",
		"--amount", "100",
	}
	err = cli.RunAdd(args, &el)

	if err != nil {
		fmt.Println(err)
	}

	err = cli.RunList(el)
	if err != nil {
		fmt.Println(err)
	}

	err = cli.RunSummary(el)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("")

	args = []string{
		"--index", "1",
		"--field", "description",
		"--description", "Ананасы",
	}
	err = cli.RunUpdate(args, &el)
	if err != nil {
		fmt.Println(err)
	}

	err = cli.RunList(el)
	if err != nil {
		fmt.Println(err)
	}

	args = []string{
		"--index", "1",
		"--field", "amount",
		"--description", "Ананасы",
		"--amount", "150",
	}

	err = cli.RunUpdate(args, &el)
	if err != nil {
		fmt.Println(err)
	}

	err = cli.RunList(el)
	if err != nil {
		fmt.Println(err)
	}
}
