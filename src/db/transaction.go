package db

import (
	"context"

	"github.com/oatsaysai/simple-core-bank/src/model"
)

type DBTransactionInterface interface {
	GetTransactionByAccountNo(accountNo string) ([]model.Transaction, error)
}

func (pgdb *PostgresqlDB) GetTransactionByAccountNo(accountNo string) ([]model.Transaction, error) {

	txs := []model.Transaction{}
	err := pgdb.DB.QueryRow(
		context.Background(),
		`
		SELECT
			COALESCE(jsonb_agg(d.*), '[]') as rows
		FROM
			(
				SELECT *
				FROM transactions
				WHERE from_account_no = $1
					OR to_account_no = $1
				ORDER BY id
			) as d
		`,
		accountNo,
	).Scan(&txs)
	if err != nil {
		return nil, err
	}

	return txs, nil
}
