package db

import (
	"database/sql"
	"errors"
	"os"
	"path"
)

type Store interface {
	DB() *sql.DB
	IsErrConstraintUnique(err error) bool
	IsErrConstraintTrigger(err error) bool
	IsErrConstraintForeignKey(err error) bool
}

func NewStore(driver string, str string) (Store, error) {
	if str == "" {
		return nil, errors.New("missing database string")
	}

	if driver == "sqlite" || driver == "sqlite3" {
		if err := os.MkdirAll(path.Dir(str), os.ModePerm); err != nil {
			return nil, err
		}
	}

	switch driver {
	case "sqlite3":
		return NewSQLite3Store(str)
	case "sqlite":
		return NewSQLiteStore(str)
	}

	return nil, errors.New(driver + " database driver is not supported")
}
