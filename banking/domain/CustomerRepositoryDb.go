package domain

import (
	"banking/customErrors"
	"banking/logger"
	"database/sql"
	"fmt"
	"os"
	"time"

	// command: go get github.com/go-sql-driver/mysql

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *customErrors.AppError) {
	var err error
	customers := make([]Customer, 0)

	// sql statement
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

		// execute Select Query findAllSql, and put the results to &customers
		err = d.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"

		// execute Select(select many rows) query findAllSql, and put the results to &customers
		// use status as query parameter
		err = d.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, customErrors.NewUnexpectedError("Unexpected DB Error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *customErrors.AppError) {
	// sql statement
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer

	// execute Get(select one row) query customerSql, and put the results to &c
	// use id as query parameter
	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customErrors.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, customErrors.NewUnexpectedError("Unexpected DB Error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)

	/**
	Function that connects to a running db instance
	This code is from https://github.com/go-sql-driver/mysql
	*/
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	// END of code from github

	return CustomerRepositoryDb{client}
}
