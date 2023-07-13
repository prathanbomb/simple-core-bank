package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createAccountTableMigration = &Migration{
	Number: 1,
	Name:   "Create account table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
			CREATE TABLE accounts(
				account_no VARCHAR(10) PRIMARY KEY NOT NULL,
				balance NUMERIC NOT NULL CHECK(balance >= 0)
			);
		`
		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create account table")
	},
}

func init() {
	Migrations = append(Migrations, createAccountTableMigration)
}
