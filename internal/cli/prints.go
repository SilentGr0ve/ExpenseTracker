package cli

import (
	"ExpenseTracker/internal/expenses"
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

func PrintHelp() {
	fmt.Println("Commands:")
	fmt.Printf("\t1. help \n")
	fmt.Printf("\t2. add {\n")
	fmt.Printf("\t3. update \n")
	fmt.Printf("\t4. delete \n")
	fmt.Printf("\t5. list \n")
	fmt.Printf("\t6. summary \n")
}

func PrintAdded(expense expenses.Expense) {
	fmt.Printf("Expense added: %s - %.2f\n", expense.Description, expense.Amount)
}

func PrintList(el expenses.ExpenseList) {
	fmt.Println("Current Expenses:")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for i, e := range el {
		fmt.Fprintf(w, "%d.\t%s\t%.2f\t%s\t\n",
			i+1, e.Description, e.Amount, e.Date.Format(time.RFC822Z),
		)
	}
	w.Flush()
}

func PrintSummary(summary float64) {
	fmt.Printf("Summary expenses: %.2f\n", summary)
}

func PrintRemove(index int) {
	fmt.Printf("\nExpense №%d removed\n", index)
}

func PrintUpdateDescription(index int, newDescription string) {
	fmt.Printf("Description updated for expense №%d: %s\n", index, newDescription)
}

func PrintUpdateAmount(index int, newAmount float64) {
	fmt.Printf("Amount updated for expense №%d: %.2f\n", index, newAmount)
}
