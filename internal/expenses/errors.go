package expenses

import "errors"

var (
	emptyDescriptionError = errors.New("description cannot be empty")
	emptyAmountError      = errors.New("amount must be greater than zero")
	emptyExpenseListError = errors.New("no expenses found")
	listOutOfRangeError   = errors.New("number of expenses are out of range")
	indexOutOfRangeError  = errors.New("index out of range")
	greaterThanZeroError  = errors.New("amount must be greater than zero")
	defaultError          = errors.New("invalid choice")
)
