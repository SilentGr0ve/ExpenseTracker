package expenses

import (
	"errors"
	"fmt"
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

func (el ExpenseList) ShowList() {
	fmt.Println("Expenses:")
	if len(el) == 0 {
		fmt.Println("No expenses found...")
		return
	}
	for _, expense := range el {
		fmt.Printf("\t%s\t%.2f\t%v\n", expense.Description, expense.Amount, expense.Date.Format(time.RFC822Z))
	}
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
