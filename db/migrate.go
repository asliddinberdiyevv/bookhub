package db

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func migrateDB(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "connecting to database")
	}

	migrationSource := "file://migrations/"
	migrator, err := migrate.NewWithDatabaseInstance(migrationSource, "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "creating migrator")
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "executing migration")
	}

	version, dirty, err := migrator.Version()
	if err != nil {
		return errors.Wrap(err, "getting migration version")
	}

	logrus.WithFields(logrus.Fields{
		"version": version,
		"dirty":   dirty,
	}).Debug("Database migrated")

	return nil
}