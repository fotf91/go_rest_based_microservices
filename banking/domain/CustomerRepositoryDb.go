package domain

import (
	"banking/customErrors"
	"database/sql"
	"log"
	"time"

	// command: go get github.com/go-sql-driver/mysql

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, *customErrors.AppError) {

	// sql statement
	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	// execute the query - Query returns multiple records
	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customer table " + err.Error())
		return nil, customErrors.NewUnexpectedError("Unexpected DB Error")
	}

	customers := make([]Customer, 0)

	// go through the rows retrieved
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

		if err != nil {
			log.Println("Error while scanning customers " + err.Error())
			return nil, customErrors.NewUnexpectedError("Unexpected DB Error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *customErrors.AppError) {
	// sql statement
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	// execute the query - QueryRow returns only one record
	row := d.client.QueryRow(customerSql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, customErrors.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, customErrors.NewUnexpectedError("Unexpected DB Error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	/**
	Function that connects to a running db instance
	This code is from https://github.com/go-sql-driver/mysql
	*/
	client, err := sql.Open("mysql", "root:codecamp@tcp(localhost:3306)/banking")
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
