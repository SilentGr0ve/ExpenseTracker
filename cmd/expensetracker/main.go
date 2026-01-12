package main

import (
	"ExpenseTracker/internal/cli"
	"ExpenseTracker/internal/storage"
	"log"
	"os"
)

const path = "storage/expenses.json"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: expensetracker <path>")
	}

	switch os.Args[1] {
	case "-h", "--help", "-help", "help":
		cli.PrintHelp()
	case "-add", "--add", "add":
		el, err := storage.LoadExpenses(path)
		if err != nil {
			log.Fatal(err)
		}
		err = cli.RunAdd(os.Args[2:], &el)
		if err != nil {
			log.Fatal(err)
		}
		err = storage.SaveExpenses(path, el)
		if err != nil {
			log.Fatal(err)
		}
	case "-summary", "--summary", "summary":
		data, err := storage.LoadExpenses(path)
		err = cli.RunSummary(data)
		if err != nil {
			log.Fatal(err)
		}
	case "-list", "--list", "list":
		data, err := storage.LoadExpenses(path)
		if err != nil {
			log.Fatal(err)
		}
		err = cli.RunList(data)
		if err != nil {
			log.Fatal(err)
		}
	case "-remove", "--remove", "remove":
		el, err := storage.LoadExpenses(path)
		if err != nil {
			log.Fatal(err)
		}
		err = cli.RunRemove(os.Args[2:], &el)
		if err != nil {
			log.Fatal(err)
		}
		err = storage.SaveExpenses(path, el)
		if err != nil {
			log.Fatal(err)
		}
	case "-update", "--update", "update":
		el, err := storage.LoadExpenses(path)
		if err != nil {
			log.Fatal(err)
		}
		err = cli.RunUpdate(os.Args[2:], &el)
		if err != nil {
			log.Fatal(err)
		}
		err = storage.SaveExpenses(path, el)
		if err != nil {
			log.Fatal(err)
		}
	}

}
