package expenses

import (
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

func (el ExpenseList) List() (ExpenseList, error) {
	if el.Len() == 0 {
		return nil, emptyExpenseListError
	}
	return el, nil
}

func (el *ExpenseList) Len() int {
	return len(*el)
}

func (el *ExpenseList) AddExpense(description string, amount float64) (*Expense, error) {
	expense, err := NewExpense(description, amount)
	if err != nil {
		return nil, err
	}
	*el = append(*el, *expense)
	return expense, nil
}

func (el *ExpenseList) RemoveExpense(index int) error {
	if index <= 0 || index > len(*el) {
		return indexOutOfRangeError
	}
	*el = slices.Delete(*el, index-1, index)
	return nil
}

func (el *ExpenseList) UpdateDescription(index int, description string) error {
	if index <= 0 || index > len(*el) {
		return indexOutOfRangeError
	}
	if description == "" {
		return emptyDescriptionError
	}
	(*el)[index-1].Description = description
	return nil
}

func (el *ExpenseList) UpdateAmount(index int, amount float64) error {
	if index <= 0 || index > len(*el) {
		return indexOutOfRangeError
	}
	if amount <= 0 {
		return greaterThanZeroError
	}
	(*el)[index-1].Amount = amount
	return nil
}

func (el *ExpenseList) Summary() (float64, error) {
	if el.Len() == 0 {
		return 0.0, emptyExpenseListError
	}
	sum := 0.0
	for _, expense := range *el {
		sum += expense.Amount
	}
	return sum, nil
}
