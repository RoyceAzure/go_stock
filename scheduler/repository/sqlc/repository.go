package repository

import "database/sql"

type Dao interface {
	Querier
}

type SQLDao struct {
	*Queries
	db *sql.DB
}

func NewSQLDao(db *sql.DB) Dao {
	return &SQLDao{
		Queries: New(db),
		db:      db,
	}
}
