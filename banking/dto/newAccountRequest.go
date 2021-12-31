package dto

import "banking/customErrors"

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *customErrors.AppError {
	if r.Amount < 5000 {
		return customErrors.NewValidationError("To open a new account you need to depose at least 5000.00")
	}
	if r.AccountType != "saving" && r.AccountType != "checking" {
		return customErrors.NewValidationError("Account type should be checking or saving")
	}
	return nil
}
