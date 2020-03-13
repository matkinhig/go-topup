package database

import (
	"database/sql"

	"github.com/matkinhig/go-topup/config"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
