package sqlitedb

import (
	"w2go/internal/app/model"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Todo struct {
	db *sqlx.DB
}

func NewTodo(db *sqlx.DB) *Todo {
	return &Todo{db: db}
}

func (ts *Todo) SelectList(q model.QueryList) ([]model.TodoDTO, error) {
	_ = sqlbuilder.NewSelectBuilder()
	return nil, nil
}

func (ts *Todo) SelectByID(id int64) {
	_ = sqlbuilder.NewSelectBuilder()
}

func (ts *Todo) DeleteByID(id []int64) {
	_ = sqlbuilder.NewDeleteBuilder()
}

func (ts *Todo) Insert() {
	_ = sqlbuilder.NewInsertBuilder()
}

func (ts *Todo) Update() {
	_ = sqlbuilder.NewUpdateBuilder()
}

func (ts *Todo) Patch() {
	_ = sqlbuilder.NewUpdateBuilder()
}
