package domain

import (
	"database/sql"
	"errors"
	"github.com/AliSayyah/banking/errs"
	"github.com/AliSayyah/banking/logger"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
	"log"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSQL string
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSQL = "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers;"
		err = d.client.Select(&customers, findAllSQL)
	} else {
		findAllSQL = "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?;"
		err = d.client.Select(&customers, findAllSQL, status)
	}
	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDB) FindByID(id int) (*Customer, *errs.AppError) {
	findByIDSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	var c Customer
	err := d.client.Get(&c, findByIDSQL, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{
		client: dbClient,
	}
}
