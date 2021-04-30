package migration

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"io/ioutil"
)

func MigrateDB(db *sqlx.DB, migrationFilePath string) error {
	buf, err := ioutil.ReadFile(migrationFilePath)
	if err != nil {
		return errors.Wrap(err, "could not read migration file")
	}

	if _, err := db.Exec(string(buf)); err != nil {
		return errors.Wrap(err, "could not execute migration file")
	}

	return nil
}
