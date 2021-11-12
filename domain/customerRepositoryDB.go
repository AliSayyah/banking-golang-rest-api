package domain

import (
	"database/sql"
	"errors"
	"github.com/AliSayyah/banking/errs"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers;"
	rows, err := d.client.Query(findAllSQL)
	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	client, err := sql.Open("mysql", GetDSN())
	if err != nil {
		panic(err)
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return CustomerRepositoryDB{client: client}
}

func GetDSN() string {
	// load .env file and get environment variables
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// get database connection string
	mysqlUser := os.Getenv("MYSQL_USER")
	MysqlPassword := os.Getenv("MYSQL_PASSWORD")
	mysqlDB := os.Getenv("MYSQL_DATABASE")
	mysqlHost := os.Getenv("MYSQL_HOST")
	mysqlPort := os.Getenv("MYSQL_PORT")

	return mysqlUser + ":" + MysqlPassword + "@tcp(" + mysqlHost + ":" + mysqlPort + ")/" + mysqlDB
}

func (d CustomerRepositoryDB) FindByID(id int) (*Customer, error) {
	findByIDSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"
	row := d.client.QueryRow(findByIDSQL, id)
	var c Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		log.Println("Error while scanning customer table " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}
	return &c, nil
}
