package sqlitedb

import (
	"context"
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

func (ts *Todo) FindMany(ctx context.Context, q model.FindManyParams) ([]model.TodoFromDB, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"id",
		"name",
		"description",
		"quantity",
		sb.As("count(*) over ()", "count"),
	)
	sb.From("todo")
	sb.Limit(q.Limit)
	sb.Offset(q.Offset)

	ApplyQueryFilters(sb, q.Filters)
	ApplyQuerySorters(sb, q.Sorters)

	sql, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	var rows []model.TodoFromDB
	err := ts.db.SelectContext(ctx, &rows, sql, args...)
	return rows, err
}

func (ts *Todo) GetOneByID(ctx context.Context, id int64) (model.TodoDTO, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
		"id",
		"name",
		"description",
		"quantity",
	)
	sb.From("todo")
	sb.Where(sb.EQ("id", id))

	sql, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	var row model.TodoDTO
	err := ts.db.GetContext(ctx, &row, sql, args...)
	return row, err
}

func (s *Todo) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	dlb := sqlbuilder.NewDeleteBuilder()
	dlb.DeleteFrom("todo")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))

	sql, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	res, err := s.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

func (ts *Todo) PatchManyByID(ctx context.Context, partials []model.TodoPartialDTO) (int64, error) {
	_ = sqlbuilder.NewUpdateBuilder()
	return 0, nil
}

func (ts *Todo) UpdateByID(ctx context.Context, todo model.TodoDTO) (int64, error) {
	_ = sqlbuilder.NewUpdateBuilder()
	return 0, nil
}

func (ts *Todo) Insert(ctx context.Context, todo model.TodoDTO) (int64, error) {
	_ = sqlbuilder.NewInsertBuilder()
	return 0, nil
}
