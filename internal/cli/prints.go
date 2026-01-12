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
	fmt.Print("\t1. 📚 help\n\n")
	fmt.Print("\t2. ✅ add -description {} -amount {};\n\t\tAdding an item with the name 'description' and price 'amount' to the list\n\n")
	fmt.Print("\t3. 🔁 update -index {} -field={} -description/-amount {}\n\t\tUpdating the value of the item in the field 'field' under the number 'index'\n\n")
	fmt.Print("\t4. ❌ delete -index {}\n\t\tRemoving item number {}\n\n")
	fmt.Print("\t5. 📋 list\n\t\tDisplay all expenses\n\n")
	fmt.Print("\t6. 💵 summary\n\t\tTotal expenses\n\n")
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
