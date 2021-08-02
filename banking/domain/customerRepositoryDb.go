package domain

import (
	"database/sql"
	"log"

	errs "banking/errors"

	"banking/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type customerRepositoryDb struct {
	db *sqlx.DB
}

func (d customerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = d.db.Query(findAllSQL)
	} else {
		findAllSQL := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = d.db.Query(findAllSQL, status)
	}
	if err != nil {
		logger.Error("error whilequerying customer table" + err.Error())
		return nil, errs.NewNotFoundError("error while querying table")
	}
	customers := make([]Customer, 0)
	err = sqlx.StructScan(rows, &customers)
	if err != nil {
		logger.Error("error while scanning customer table" + err.Error())
		return nil, errs.UnexpectedDatabaseError("error  while scanning")
	}
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Dateofbirth, &c.Status)
	// 	if err != nil {
	// 		log.Println("error while scanning customer table" + err.Error())
	// 		return nil, errs.UnexpectedDatabaseError("error  while scanning")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil

}

func (d customerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	customerSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	rows := d.db.QueryRow(customerSql, id)

	var c Customer
	err := rows.Scan(&c.Id, &c.Name, &c.Dateofbirth, &c.Zipcode, &c.City, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("error while scanning customers table" + err.Error())
			return nil, errs.UnexpectedDatabaseError("databse error")

		}

	}
	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) customerRepositoryDb {

	return customerRepositoryDb{db: dbClient}

}
