package domain

import (
	errs "banking/errors"
	"banking/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	db *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT into accounts(customer_id, opening_date, account_type, amount, status) values (?,?,?,?,?)"
	result, err := d.db.Exec(sqlInsert, a.CustomerID, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("error while creating new account" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("error while getting last inserted account" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{db: dbClient}
}
