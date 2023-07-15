package db

import (
	"context"
	"database/sql"
	"github.com/shopspring/decimal"
)

type DBAccountInterface interface {
	InsertAccount(accountNo, accountName string, balance decimal.Decimal) error
	GetAccount(accountNo string) (*string, *string, *decimal.Decimal, error)
	AccountExists(accountNo string) (bool, error)
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

func (pgdb *PostgresqlDB) AccountExists(accountNo string) (bool, error) {
	var exists bool

	query := `
		SELECT EXISTS (
			SELECT 1 
			FROM accounts
			WHERE account_no = $1
		);
	`

	err := pgdb.DB.QueryRow(context.Background(), query, accountNo).Scan(&exists)
	if err != nil {
		pgdb.logger.Errorf("Failed to check if account exists with account no %s: %v", accountNo, err)
		return false, err
	}

	if exists {
		pgdb.logger.Infof("Account with account no %s exists", accountNo)
	} else {
		pgdb.logger.Infof("Account with account no %s does not exist", accountNo)
	}

	return exists, nil
}

func (pgdb *PostgresqlDB) GetAccount(accountNo string) (*string, *string, *decimal.Decimal, error) {
	var accountName string
	var balance decimal.Decimal

	query := `
		SELECT account_name, balance 
		FROM accounts
		WHERE account_no = $1;
	`

	err := pgdb.DB.QueryRow(context.Background(), query, accountNo).Scan(&accountName, &balance)
	if err != nil {
		if err == sql.ErrNoRows {
			pgdb.logger.Errorf("No account found with account number %s", accountNo)
			return nil, nil, nil, err
		}
		pgdb.logger.Errorf("Failed to get account details for account no %s: %v", accountNo, err)
		return nil, nil, nil, err
	}

	pgdb.logger.Infof("Retrieved account details for account no %s", accountNo)

	return &accountNo, &accountName, &balance, nil
}
