package sqlitedb

import (
	"context"
	"fmt"

	"gostart-crm/internal/app/model"

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

	ApplyQueryLimitOffset(sb, q.Limit, q.Offset)
	ApplyQueryFilters(sb, q.Filters)
	ApplyQuerySorters(sb, q.Sorters)

	sql, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	var rows []model.TodoFromDB
	return rows, st.db.SelectContext(ctx, &rows, sql, args...)
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
	return row, st.db.GetContext(ctx, &row, sql, args...)
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
		updates[i] = ub

		if p.Quantity != nil {
			if value, ok := (*p.Quantity).(float64); ok {
				ub.SetMore(ub.EQ("quantity", value))
			} else {
				ub.SetMore(ub.EQ("quantity", nil))
			}
		}
	}

	tx, err := st.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	var affectedTotal int64

	for _, ub := range updates {
		sql, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
		fmt.Println(sql, args)

		res, err := tx.ExecContext(ctx, sql, args...)
		if err != nil {
			return 0, err
		} else {
			affectedUpdate, _ := res.RowsAffected()
			affectedTotal += affectedUpdate
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	} else {
		return affectedTotal, nil
	}
}

func (st *Todo) UpdateByID(ctx context.Context, todo model.TodoDTO) (int64, error) {
	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("todo")
	ub.Where(ub.EQ("id", todo.ID))
	ub.Set(
		ub.EQ("name", todo.Name),
		ub.EQ("description", todo.Description),
		ub.EQ("quantity", todo.Quantity),
	)

	sql, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	res, err := st.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func (st *Todo) Insert(ctx context.Context, todo model.TodoDTO) (int64, error) {
	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("todo")
	ib.Cols("name", "description", "quantity")
	ib.Values(todo.Name, todo.Description, todo.Quantity)

	sql, args := ib.BuildWithFlavor(sqlbuilder.SQLite)
	fmt.Println(sql, args)

	res, err := st.db.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}
