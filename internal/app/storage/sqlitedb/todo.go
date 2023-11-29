package sqlitedb

import (
	"fmt"

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

func (ts *Todo) SelectList(p model.QueryListParams) ([]model.TodoDTO, int64, error) {
	sb := sqlbuilder.NewSelectBuilder().
		Select("*").
		From("todo").
		Limit(p.Limit).
		Offset(p.Offset)

	applyQueryWhere(sb, p.Where)
	applyQueryOrderBy(sb, p.OrderBy)

	sql, args := sb.Build()
	fmt.Println(sql, args)
	return nil, 0, nil
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
