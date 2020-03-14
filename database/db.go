package database

import (
	"github.com/jmoiron/sqlx"
	"github.com/matkinhig/go-topup/config"
)

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
