package db

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/shopspring/decimal"
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
		latestNumber, err = strconv.Atoi(latestAccountNo.String)
		if err != nil {
			logger.Errorf("Failed to parse latest account number: %+v", err)
			return fmt.Errorf("failed to parse latest account number: %w", err)
		}
	} else {
		// This is the first time running the generation operation
		logger.Info("First time generating account numbers")
	}

	// Generate new account numbers
	rows := [][]interface{}{}
	for i := 0; i < batchSize; i++ {
		rows = append(rows, []interface{}{
			fmt.Sprintf("%010d", latestNumber+i+1),
		})
	}

	// Shuffle account numbers
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rows), func(i, j int) { rows[i], rows[j] = rows[j], rows[i] })

	// Execute batch insert operation
	_, err = pgdb.DB.CopyFrom(
		context.Background(),
		pgx.Identifier{"pregen_acc_no"},
		[]string{"account_no"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		logger.Errorf("Failed to execute CopyFrom operation: %+v", err)
		return fmt.Errorf("failed to execute CopyFrom operation: %w", err)
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
	defer tx.Rollback(context.Background())

	err = tx.QueryRow(
		context.Background(),
		"SELECT account_no FROM pregen_acc_no WHERE id = (SELECT nextval('available_acc_no_id'))",
	).Scan(&accountNo)
	if err != nil {
		logger.Errorf("Failed to scan account number: %+v", err)
		return "", fmt.Errorf("failed to scan account number: %w", err)
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
