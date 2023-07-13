package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createTransactionTableMigration = &Migration{
	Number: 2,
	Name:   "Create transaction table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
			CREATE TABLE transactions(
				id BIGSERIAL PRIMARY KEY NOT NULL,
				from_account_no VARCHAR(10), -- Must be null if tx is transfer-in
				to_account_no VARCHAR(10),   -- Must be null if tx is transfer-out
				amount NUMERIC NOT NULL,
				transaction_type VARCHAR(12) NOT NULL,
				transaction_time TIMESTAMPTZ NOT NULL DEFAULT NOW()
			);
		`
		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create transaction table")
	},
}

func init() {
	Migrations = append(Migrations, createTransactionTableMigration)
}
