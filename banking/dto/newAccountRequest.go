package dto

import (
	errs "banking/errors"
	"fmt"
	"strings"
)

type NewAccountRequest struct {
	CustomerID  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("to open a new account deposit 500")
	}
	fmt.Println(r.AccountType)
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("account type should be checking or saving")

	}
	return nil
}
