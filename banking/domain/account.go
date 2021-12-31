package domain

import (
	"banking/customErrors"
	"banking/dto"
)

const dbTSLayout = "2006-01-02 15:04:05"

type Account struct {
	AccountId   string `db:"account_id"`
	CustomerId  string `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      string
}

func (a Account) CanWithdraw(amount float64) bool {
	if a.Amount < amount {
		return false
	}
	return true
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

type AccountRepository interface {
	Save(Account) (*Account, *customErrors.AppError)
	FindBy(string) (*Account, *customErrors.AppError)
	SaveTransaction(Transaction) (*Transaction, *customErrors.AppError)
}
