package expenses

import "errors"

var emptyDescriptionError = errors.New("description cannot be empty")
var emptyAmountError = errors.New("amount must be greater than zero")
var emptyExpenseListError = errors.New("no expenses found")
var listOutOfRangeError = errors.New("number of expenses are out of range")
var indexOutOfRangeError = errors.New("index out of range")
var greaterThanZeroError = errors.New("amount must be greater than zero")
var defaultError = errors.New("invalid choice")
