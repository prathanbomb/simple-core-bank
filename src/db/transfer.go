package db

import (
	"context"

	"github.com/oatsaysai/simple-core-bank/src/custom_error"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type TransactionType string

const (
	TRANSFER_IN  TransactionType = "transfer-in"
	TRANSFER     TransactionType = "transfer"
	TRANSFER_OUT TransactionType = "transfer-out"
)

type DBTransferInterface interface {
	TransferIn(toAccountNo string, amount decimal.Decimal) (*int64, error)
	TransferOut(fromAccountNo string, amount decimal.Decimal) (*int64, error)
}

func (pgdb *PostgresqlDB) TransferIn(toAccountNo string, amount decimal.Decimal) (*int64, error) {
	logger := pgdb.logger

	tx, err := pgdb.DB.Begin(context.Background())
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, errors.Wrap(err, "Unable to make a transaction")
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		`
			SET TRANSACTION ISOLATION LEVEL READ COMMITTED
		`,
	)
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, err
	}

	cmdTag, err := tx.Exec(
		context.Background(),
		`
			UPDATE accounts
				SET balance = balance + $2
			WHERE account_no = $1;

		`,
		toAccountNo,
		amount,
	)
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, err
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, &custom_error.UserError{
			Code:    custom_error.AccountNoNotFound,
			Message: "to_account_no not found",
		}
	}

	var transactionID int64
	tx.QueryRow(
		context.Background(),
		`
			INSERT INTO transactions(
				to_account_no,
				amount,
				transaction_type
			)
		 	VALUES ($1, $2, $3)
			RETURNING id;
		`,
		toAccountNo,
		amount,
		TRANSFER_IN,
	).Scan(&transactionID)

	err = tx.Commit(context.Background())
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, errors.Wrap(err, "Unable to commit a transaction")
	}

	return &transactionID, nil
}

func (pgdb *PostgresqlDB) TransferOut(fromAccountNo string, amount decimal.Decimal) (*int64, error) {
	logger := pgdb.logger

	tx, err := pgdb.DB.Begin(context.Background())
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, errors.Wrap(err, "Unable to make a transaction")
	}
	defer tx.Rollback(context.Background())

	_, err = tx.Exec(
		context.Background(),
		`
			SET TRANSACTION ISOLATION LEVEL READ COMMITTED
		`,
	)
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, err
	}

	cmdTag, err := tx.Exec(
		context.Background(),
		`
			UPDATE accounts
				SET balance = balance - $2
			WHERE account_no = $1;

		`,
		fromAccountNo,
		amount,
	)
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, &custom_error.UserError{
			Code:    custom_error.BalanceNotEnough,
			Message: "account balance not enough",
		}
	}
	if cmdTag.RowsAffected() == 0 {
		return nil, &custom_error.UserError{
			Code:    custom_error.AccountNoNotFound,
			Message: "from_account_no not found",
		}
	}

	var transactionID int64
	tx.QueryRow(
		context.Background(),
		`
			INSERT INTO transactions(
				from_account_no,
				amount,
				transaction_type
			)
		 	VALUES ($1, $2, $3)
			RETURNING id;
		`,
		fromAccountNo,
		amount,
		TRANSFER_OUT,
	).Scan(&transactionID)

	err = tx.Commit(context.Background())
	if err != nil {
		logger.Errorf("%+v", err)
		return nil, errors.Wrap(err, "Unable to commit a transaction")
	}

	return &transactionID, nil
}
