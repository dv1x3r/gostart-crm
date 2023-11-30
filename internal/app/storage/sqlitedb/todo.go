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

func (st *Todo) FindMany(ctx context.Context, q model.FindManyParams) ([]model.TodoFromDB, error) {
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
	err := st.db.SelectContext(ctx, &rows, sql, args...)
	return rows, err
}

func (st *Todo) GetOneByID(ctx context.Context, id int64) (model.TodoDTO, error) {
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
	err := st.db.GetContext(ctx, &row, sql, args...)
	return row, err
}

func (st *Todo) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	dlb := sqlbuilder.NewDeleteBuilder()
	dlb.DeleteFrom("todo")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))

	sql, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	res, err := st.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func (st *Todo) PatchManyByID(ctx context.Context, partials []model.TodoPartialDTO) (int64, error) {
	updates := make([]*sqlbuilder.UpdateBuilder, len(partials))

	for i, p := range partials {
		ub := sqlbuilder.NewUpdateBuilder()
		ub.Update("todo")
		ub.Where(ub.EQ("id", p.ID))

		if p.Quantity != nil {
			if value, ok := (*p.Quantity).(float64); ok {
				ub.Set(ub.EQ("quantity", value))
			} else {
				ub.Set(ub.EQ("quantity", nil))
			}
		}

		updates[i] = ub
	}

	tx, err := st.db.Begin()
	if err != nil {
		return 0, err
	}

	var affected int64

	for _, ub := range updates {
		sql, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
		fmt.Println(sql, args)

		res, err := tx.ExecContext(ctx, sql, args...)
		if err != nil {
			return 0, tx.Rollback()
		} else {
			ra, _ := res.RowsAffected()
			affected += ra
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	} else {
		return affected, nil
	}
}

func (st *Todo) UpdateByID(ctx context.Context, todo model.TodoDTO) (int64, error) {
	_ = sqlbuilder.NewUpdateBuilder()
	return 0, nil
}

func (st *Todo) Insert(ctx context.Context, todo model.TodoDTO) (int64, error) {
	_ = sqlbuilder.NewInsertBuilder()
	return 0, nil
}
