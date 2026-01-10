package cli

import (
	"ExpenseTracker/internal/expenses"
)

func RunAdd(args []string, el *expenses.ExpenseList) error {

	flags, err := ParseAddFlags(args)
	if err != nil {
		return err
	}

	expense, err := el.AddExpense(flags.Description, flags.Amount)
	if err != nil {
		return err
	}

	PrintAdded(*expense)
	return nil
}

func RunList(el expenses.ExpenseList) error {
	el, err := el.List()
	if err != nil {
		return err
	}
	PrintList(el)
	return nil
}

func RunSummary(el expenses.ExpenseList) error {
	summary, err := el.Summary()
	if err != nil {
		return err
	}
	PrintSummary(summary)
	return nil
}

func RunRemove(args []string, el *expenses.ExpenseList) error {
	flags, err := ParseRemoveFlags(args, el.Len())
	if err != nil {
		return err
	}

	err = el.RemoveExpense(flags.Index)
	if err != nil {
		return err
	}

	PrintRemove(flags.Index)
	return nil
}

func RunUpdate(args []string, el *expenses.ExpenseList) error {
	flags, err := ParseUpdateFlags(args, el.Len())
	if err != nil {
		return err
	}

	if flags.Field == "description" {
		err = el.UpdateDescription(flags.Index, flags.Description)
		if err != nil {
			return err
		}
		PrintUpdateDescription(flags.Index, flags.Description)
	} else if flags.Field == "amount" {
		err = el.UpdateAmount(flags.Index, flags.Amount)
		if err != nil {
			return err
		}
		PrintUpdateAmount(flags.Index, flags.Amount)
	}
	return nil
}
