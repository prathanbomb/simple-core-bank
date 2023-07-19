package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var createAccountNumbersTableMigration = &Migration{
	Number: 3,
	Name:   "Create AccountNumbers table",
	Forwards: func(db *gorm.DB) error {
		const sql = `
			CREATE TABLE pregen_acc_no(
				account_no VARCHAR(10) PRIMARY KEY NOT NULL,
				is_used BOOLEAN NOT NULL DEFAULT FALSE
			);
		`
		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create AccountNumbers table")
	},
}

func init() {
	Migrations = append(Migrations, createAccountNumbersTableMigration)
}
