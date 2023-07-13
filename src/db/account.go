package db

import (
	"context"

	"github.com/shopspring/decimal"
)

type DBAccountInterface interface {
	InsertAccount(accountNo, accountName string, balance decimal.Decimal) error
}

func (pgdb *PostgresqlDB) InsertAccount(accountNo, accountName string, balance decimal.Decimal) error {
	logger := pgdb.logger

	_, err := pgdb.DB.Exec(context.Background(),
		`
		INSERT INTO accounts(
			account_no,
			account_name,
			balance
		)
		VALUES ($1, $2, $3);
  		`,
		accountNo,
		accountName,
		balance,
	)
	if err != nil {
		logger.Errorf("%+v", err)
		return err
	}

	return nil
}
