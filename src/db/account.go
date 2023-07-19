package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type DBAccountInterface interface {
	GetAccount(accountNo string) (*string, *string, *decimal.Decimal, error)
	AccountExists(accountNo string) (bool, error)
	GetAccountNoAndInsertAccount(accountName string, balance decimal.Decimal) (string, error)
	PreGenerateAccountNo(batchSize int) error
}

func (pgdb *PostgresqlDB) PreGenerateAccountNo(batchSize int) error {
	logger := pgdb.logger

	// Query the largest account number
	var latestAccountNo sql.NullString
	err := pgdb.DB.QueryRow(context.Background(), `SELECT MAX(account_no) FROM pregen_acc_no`).Scan(&latestAccountNo)

	// Initialize latestNumber to 0
	latestNumber := 0

	if err != nil {
		logger.Errorf("Failed to query latest account number: %+v", err)
		return fmt.Errorf("failed to query latest account number: %w", err)
	}

	if latestAccountNo.Valid {
		// Parse the numeric part of the latest account number
		latestNumber, err = strconv.Atoi(latestAccountNo.String[3:])
		if err != nil {
			logger.Errorf("Failed to parse latest account number: %+v", err)
			return fmt.Errorf("failed to parse latest account number: %w", err)
		}
	} else {
		// This is the first time running the generation operation
		logger.Info("First time generating account numbers")
	}

	// Generate new account numbers and build insert values string
	values := make([]string, batchSize)
	for i := range values {
		accountNo := fmt.Sprintf("%s%07d", "007", latestNumber+i+1)
		values[i] = fmt.Sprintf("('%s', false)", accountNo)
	}

	// Execute batch insert operation
	_, err = pgdb.DB.Exec(context.Background(), `INSERT INTO pregen_acc_no(account_no, is_used) VALUES `+strings.Join(values, ","))
	if err != nil {
		logger.Errorf("Failed to execute batch insert operation: %+v", err)
		return fmt.Errorf("failed to execute batch insert operation: %w", err)
	}

	logger.Infof("Successfully pre-generated %d account numbers", batchSize)
	return nil
}

func (pgdb *PostgresqlDB) GetAccountNoAndInsertAccount(accountName string, balance decimal.Decimal) (string, error) {
	var accountNo string
	logger := pgdb.logger

	tx, err := pgdb.DB.Begin(context.Background())
	if err != nil {
		logger.Errorf("Failed to begin transaction: %+v", err)
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if p := recover(); p != nil {
			if err := tx.Rollback(context.Background()); err != nil {
				logger.Errorf("Failed to rollback transaction after panic: %+v", err)
			}
			logger.Errorf("Transaction rolled back due to a panic: %+v", p)
		}
	}()

	row := tx.QueryRow(context.Background(), "SELECT account_no FROM pregen_acc_no WHERE is_used = FALSE LIMIT 1 FOR UPDATE")
	if err = row.Scan(&accountNo); err != nil {
		if rbErr := tx.Rollback(context.Background()); rbErr != nil {
			logger.Errorf("Failed to rollback transaction after failing to scan account number: %+v", rbErr)
		}
		logger.Errorf("Failed to scan account number: %+v", err)
		return "", fmt.Errorf("failed to scan account number: %w", err)
	}

	if _, err = tx.Exec(context.Background(), "UPDATE pregen_acc_no SET is_used = TRUE WHERE account_no = $1", accountNo); err != nil {
		if rbErr := tx.Rollback(context.Background()); rbErr != nil {
			logger.Errorf("Failed to rollback transaction after failing to execute update: %+v", rbErr)
		}
		logger.Errorf("Failed to execute update on account number: %+v", err)
		return "", fmt.Errorf("failed to execute update on account number: %w", err)
	}

	_, err = tx.Exec(context.Background(),
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
		if rbErr := tx.Rollback(context.Background()); rbErr != nil {
			logger.Errorf("Failed to rollback transaction after failing to execute insert: %+v", rbErr)
		}
		logger.Errorf("Failed to insert account: %+v", err)
		return "", fmt.Errorf("failed to insert account: %w", err)
	}

	if err = tx.Commit(context.Background()); err != nil {
		logger.Errorf("Failed to commit transaction: %+v", err)
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	logger.Infof("Successfully got and marked account number as used: %s", accountNo)
	return accountNo, nil
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
