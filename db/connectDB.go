package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func ConnectDB() *sqlx.DB {
	db_driver := os.Getenv("DB_DRIVER")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", db_driver, db_user, db_password, db_host, db_port, db_name)

	conn, err := sqlx.Open(db_driver, connStr)
	if err != nil {
		logrus.Fatalf("failed connecting to database error: %s", err.Error())
		os.Exit(1)
	}

	conn.SetMaxOpenConns(32)

	pingErr := conn.Ping()
	logrus.Infof("db ping status: %v", pingErr)
	if pingErr != nil {
		logrus.Fatalf("failed connecting to db error: %s", pingErr.Error())
		os.Exit(1)
	}

	logrus.Infof("successfully connected to db")

	return conn
}
