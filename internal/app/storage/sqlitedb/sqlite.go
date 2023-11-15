package sqlitedb

import "github.com/jmoiron/sqlx"

type Todo struct {
	db *sqlx.DB
}

func NewTodo(db *sqlx.DB) *Todo {
	return &Todo{db: db}
}
