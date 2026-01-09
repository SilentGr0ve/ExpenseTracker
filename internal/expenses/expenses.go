package expenses

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"slices"
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
		return nil, emptyDescriptionError
	}
	if amount <= 0 {
		return nil, emptyAmountError
	}
	return &Expense{
		Description: description,
		Amount:      amount,
		Date:        time.Now(),
	}, nil
}

//func (el *ExpenseList) List() (ExpenseList, error) {
//	fmt.Println("Current expenses:")
//	if len(*el) == 0 {
//		return ExpenseList{}, emptyExpenseListError
//	}
//	return *el, nil
//}

func (el ExpenseList) ShowList() error {
	fmt.Println("Current expenses:")
	if len(el) == 0 {
		return emptyExpenseListError
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
	_, err = fmt.Scanf("%d", &count)
	if err != nil {
		return NaNError
	}
	if count == 0 || count > len(*el) {
		return listOutOfRangeError
	}

	fmt.Printf("\nExpense в„–%d removed\n", count)
	*el = slices.Delete(*el, count-1, (count-1)+1)
	return nil
}

func (el *ExpenseList) Update() error {
	err := el.ShowList()
	if err != nil {
		return err
	}

	fmt.Printf("\nEnter number of expense to update: ")
	var count int
	_, err = fmt.Scanf("%d", &count)
	if err != nil {
		return NaNError
	}
	if count == 0 || count > len(*el) {
		return listOutOfRangeError
	}

	fmt.Printf("What exactly do you want to update?\n\t1. Description\n\t2. Amount\n")
	var choice int
	_, err = fmt.Scanf("%d", &choice)
	if err != nil {
		return NaNError
	}
	switch choice {
	case 1:
		fmt.Printf("\nEnter the new description of the expense: ")
		var description string
		_, _ = fmt.Scanf("%s", &description)
		err := el.UpdateDescription(count, description)
		if err != nil {
			return err
		}

	case 2:
		fmt.Printf("\nEnter the new amount of the expense: ")
		var amount float64
		_, _ = fmt.Scanf("%f", &amount)
		err := el.UpdateAmount(count, amount)
		if err != nil {
			return err
		}

	default:
		return defaultError
	}
	return nil
}

func (el *ExpenseList) UpdateDescription(index int, description string) error {
	if description == "" {
		return errors.New("description cannot be empty")
	}
	if index <= 0 || index > len(*el) {
		return indexOutOfRangeError
	}
	(*el)[index].Description = description
	fmt.Printf("Description updated: %s\n", (*el)[index].Description)
	return nil
}

func (el *ExpenseList) UpdateAmount(index int, amount float64) error {
	if amount <= 0 {
		return greaterThanZeroError
	}
	if index <= 0 || index > len(*el) {
		return indexOutOfRangeError
	}
	(*el)[index].Amount = amount
	fmt.Printf("Amount updated: %.2f\n", (*el)[index].Amount)
	return nil
}

func (el *ExpenseList) Summary() (float64, error) {
	if len(*el) == 0 {
		return 0.0, emptyExpenseListError
	}
	sum := 0.0
	for _, expense := range *el {
		sum += expense.Amount
	}
	return sum, nil
}
