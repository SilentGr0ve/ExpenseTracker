package main

import (
	"ExpenseTracker/internal/cli"
	"ExpenseTracker/internal/expenses"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	el := expenses.NewExpenseList()

	for {
		fmt.Print("> ")
		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)

			switch input {
			case "help":
				cli.RunHelp()
			case "add":
				err := cli.RunAdd(os.Args[2:], &el)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			case "list":
				//
			default:
				fmt.Println("Error: Invalid input!")
				os.Exit(1)
			}
		}
	}

}
