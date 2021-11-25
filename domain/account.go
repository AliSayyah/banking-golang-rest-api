package domain

import (
	"github.com/AliSayyah/banking/dto"
	"github.com/AliSayyah/banking/errs"
)

type Account struct {
	AccountID   int
	CustomerID  int
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToAccountResponseDTO() dto.AccountResponse {
	return dto.AccountResponse{
		AccountID: a.AccountID,
	}
}

type AccountRepository interface {
	Save(account Account) (*Account, *errs.AppError)
}
