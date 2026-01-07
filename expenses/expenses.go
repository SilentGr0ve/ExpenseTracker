package expenses

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"text/tabwriter"
	"time"
)

type Expense struct {
	Description string
	Amount      float64
	Date        time.Time
}

type ExpenseList []Expense

func NewExpenseList() ExpenseList {
	return make(ExpenseList, 0)
}

func NewExpense(description string, amount float64) (*Expense, error) {
	if description == "" {
		return nil, errors.New("description cannot be empty")
	}
	if amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}
	return &Expense{
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}, nil
}

func (el ExpenseList) ShowList() error {
	fmt.Println("Current expenses:")
	if len(el) == 0 {
		return errors.New("no expenses found")
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	for i, e := range el {
		fmt.Fprintf(w, "%d.\t%s\t%.2f\t%s\t\n",
			i+1, e.Description, e.Amount, e.Date.Format(time.RFC822Z),
		)
	}
	w.Flush()
	return nil
}

func (el *ExpenseList) AddExpense(description string, amount float64) error {
	expense, err := NewExpense(description, amount)
	if err != nil {
		return err
	}
	*el = append(*el, *expense)
	fmt.Printf("Expense added: %s - %.2f\n", expense.Description, expense.Amount)
	return nil
}

func (el *ExpenseList) RemoveExpense() error {
	err := el.ShowList()
	if err != nil {
		return err
	}

	fmt.Printf("\nEnter number of expense to remove: ")
	var count int
	_, _ = fmt.Scanf("%d", &count)
	if count == 0 || count > len(*el) {
		return errors.New("number of expenses are out of range")
	}

	fmt.Printf("\nExpense №%d removed\n", count)
	*el = slices.Delete(*el, count-1, (count-1)+1)
	return nil
}

func (el *ExpenseList) UpdateExpense() error {
	err := el.ShowList()
	if err != nil {
		return err
	}

	fmt.Printf("\nEnter number of expense to update: ")
	var count int
	_, _ = fmt.Scanf("%d", &count)
	if count == 0 || count > len(*el) {
		return errors.New("number of expenses are out of range")
	}

	fmt.Printf("What exactly do you want to update?\n\t1. Description\n\t2. Amount\n\t3. Both of them\n")
	return nil
	// Не закончено
}

func (el *ExpenseList) Summary() (float64, error) {
	if len(*el) == 0 {
		return 0.0, errors.New("no expenses found")
	}
	sum := 0.0
	for _, expense := range *el {
		sum += expense.Amount
	}
	return sum, nil
}
