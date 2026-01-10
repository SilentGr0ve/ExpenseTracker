package cli

import (
	"errors"
	"flag"
)

type AddFlags struct {
	Description string
	Amount      float64
}

func ParseAddFlags(args []string) (*AddFlags, error) {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	desc := fs.String("description", "", "expense description")
	amount := fs.Float64("amount", 0, "amount of expense")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if *desc == "" {
		return nil, errors.New("description cannot be empty")
	}
	if *amount <= 0 {
		return nil, errors.New("amount must be greater than zero")
	}

	return &AddFlags{
		Description: *desc,
		Amount:      *amount,
	}, nil
}

type RemoveFlags struct {
	Index int
}

func ParseRemoveFlags(args []string, length int) (*RemoveFlags, error) {
	fs := flag.NewFlagSet("remove", flag.ContinueOnError)
	index := fs.Int("index", 0, "expense index")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if *index <= 0 || *index > length {
		return nil, errors.New("index out of range")
	}

	return &RemoveFlags{
		Index: *index,
	}, nil
}

type UpdateFlags struct {
	Index       int
	Field       string
	Description string
	Amount      float64
}

func ParseUpdateFlags(args []string, length int) (*UpdateFlags, error) {
	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	index := fs.Int("index", 0, "expense index")
	field := fs.String("field", "", "expense field name (description or amount)")
	description := fs.String("description", "", "new expense description")
	amount := fs.Float64("amount", 0, "new amount of expense")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if *index <= 0 || *index > length {
		return nil, errors.New("index out of range")
	}

	if *field != "description" && *field != "amount" {
		return nil, errors.New("field must be 'description' or 'amount'")
	}

	if *field == "description" {
		if *description == "" {
			return nil, errors.New("description cannot be empty")
		}
	} else if *field == "amount" {
		if *amount <= 0 {
			return nil, errors.New("amount must be greater than zero")
		}
	}

	return &UpdateFlags{
		Index:       *index,
		Field:       *field,
		Description: *description,
		Amount:      *amount,
	}, nil
}
