package db

import "database/sql"

type Database struct {
	*sql.DB
}

func NewDataBase(db *sql.DB) *Database {
	return &Database{
		db,
	}
}
