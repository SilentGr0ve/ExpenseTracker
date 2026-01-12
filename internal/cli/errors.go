package cli

import "errors"

var (
	emptyAmountError      = errors.New("amount must be greater than zero")
	emptyDescriptionError = errors.New("description cannot be empty")
	indexOutOfRangeError  = errors.New("index out of range")
	fieldValueError       = errors.New("field value must be 'description' or 'amount'")
)
