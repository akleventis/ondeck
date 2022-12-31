package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

type DB struct {
	*sql.DB
}

func Open() (*DB, error) {
	var (
		dbUser = os.Getenv("DB_USER")
		dbName = os.Getenv("DB_NAME")
		dbHost = os.Getenv("DB_HOST")
		dbPort = os.Getenv("DB_PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	logrus.Printf("%s database connected", dbName)

	odDB := &DB{
		db,
	}

	return odDB, nil
}
