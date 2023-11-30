package sqlitedb

import (
	"context"
	"database/sql"
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

type TodoBase struct {
	ID          int64          `db:"id"`
	Name        string         `db:"name"`
	Description sql.NullString `db:"description"`
	Quantity    sql.NullInt64  `db:"quantity"`
}

type TodoDetails struct {
	TodoBase
	Total int64 `db:"total"`
}

func (ts *Todo) SelectList(q model.QuerySelectList) ([]model.TodoDTO, int64, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb = sb.
		Select(
			"id",
			"name",
			"description",
			"quantity",
			sb.As("count(*) over ()", "total"),
		).
		From("todo").
		Limit(q.Limit).
		Offset(q.Offset)

	ApplyQueryFilters(sb, q.Filters)
	ApplyQuerySorters(sb, q.Sorters)

	sql, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	var rows []TodoDetails
	err := ts.db.SelectContext(context.TODO(), &rows, sql, args...)

	todos := make([]model.TodoDTO, len(rows))
	var total int64
	for i, row := range rows {
		total = row.Total
		todos[i] = model.TodoDTO{
			ID:          row.ID,
			Name:        row.Name,
			Description: &row.Description.String,
			Quantity:    &row.Quantity.Int64,
		}
	}

	return todos, total, err
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
