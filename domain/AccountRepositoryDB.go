package domain

import (
	"github.com/AliSayyah/banking/errs"
	"github.com/AliSayyah/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)`
	result, err := d.client.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting new account's ID: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountID = int(id)
	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{client: dbClient}
}
