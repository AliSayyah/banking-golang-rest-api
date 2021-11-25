package service

import (
	"github.com/AliSayyah/banking/domain"
	"github.com/AliSayyah/banking/dto"
	"github.com/AliSayyah/banking/errs"
	"time"
)

type AccountService interface {
	NewAccount(request dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(r dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {
	var err = r.Validate()
	if err != nil {
		return nil, err
	}
	a := domain.Account{
		AccountID:   0,
		CustomerID:  r.CustomerID,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: r.AccountType,
		Amount:      r.Amount,
		Status:      "1",
	}
	newAccount, err := s.repo.Save(a)
	if err != nil {
		return nil, err
	}
	response := newAccount.ToAccountResponseDTO()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repo,
	}
}
