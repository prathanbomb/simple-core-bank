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
				id BIGSERIAL PRIMARY KEY NOT NULL,
				account_no VARCHAR(10) NOT NULL
			);
			CREATE SEQUENCE available_acc_no_id START 1;
		`
		err := db.Exec(sql).Error
		return errors.Wrap(err, "unable to create AccountNumbers table")
	},
}

func init() {
	Migrations = append(Migrations, createAccountNumbersTableMigration)
}
