package sqlitedb

import (
	"context"
	"database/sql"
	"errors"

	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/jmoiron/sqlx"
	"github.com/mattn/go-sqlite3"
)

type SelectProvider interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

func logDebugQuery(ctx context.Context, query string, args any) {
	if !utils.GetConfig().DebugSQL {
		return
	}

	log := utils.GetLogger()
	log.Debug().Any("id", ctx.Value("RequestID")).
		Str("sql", query).
		Any("sql_args", args).
		Msg("stor")
}

func runSelect[T any](ctx context.Context, sp SelectProvider, query string, args []any) ([]T, error) {
	logDebugQuery(ctx, query, args)
	var rows []T
	return rows, sp.SelectContext(ctx, &rows, query, args...)
}

func runGet[T any](ctx context.Context, sp SelectProvider, query string, args []any) (T, bool, error) {
	logDebugQuery(ctx, query, args)
	var row T
	err := sp.GetContext(ctx, &row, query, args...)
	if errors.Is(err, sql.ErrNoRows) {
		return row, false, nil
	} else if err != nil {
		return row, false, err
	}
	return row, true, nil
}

// "github.com/mattn/go-sqlite3"
func runExecResult(ctx context.Context, ec sqlx.ExecerContext, query string, args []any) (sql.Result, error) {
	res, err := ec.ExecContext(ctx, query, args...)

	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.ExtendedCode == 2067 {
			err = errors.Join(storage.ErrConstraintUnique, err)
		} else if sqliteErr.ExtendedCode == 1811 {
			err = errors.Join(storage.ErrConstraintTrigger, err)
		} else if sqliteErr.ExtendedCode == 787 {
			err = errors.Join(storage.ErrConstraintForeignKey, err)
		}
	}

	return res, err
}

func runExecAffected(ctx context.Context, ec sqlx.ExecerContext, query string, args []any) (int64, error) {
	logDebugQuery(ctx, query, args)
	if res, err := runExecResult(ctx, ec, query, args); err != nil {
		return 0, err
	} else {
		return res.RowsAffected()
	}
}

func runExecInsert(ctx context.Context, ec sqlx.ExecerContext, query string, args []any) (int64, error) {
	logDebugQuery(ctx, query, args)
	if res, err := runExecResult(ctx, ec, query, args); err != nil {
		return 0, err
	} else {
		return res.LastInsertId()
	}
}

func runExecAffectedRangeTx(ctx context.Context, tx *sqlx.Tx, queries []string, args [][]any) (int64, error) {
	var total int64
	for i := range queries {
		affected, err := runExecAffected(ctx, tx, queries[i], args[i])
		if err != nil {
			return 0, err
		}
		total += affected
	}
	return total, nil
}

func runExecAffectedRangeNewTx(ctx context.Context, db *sqlx.DB, queries []string, args [][]any) (int64, error) {
	tx, err := db.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	affected, err := runExecAffectedRangeTx(ctx, tx, queries, args)
	if err != nil {
		return 0, err
	}

	return affected, tx.Commit()
}
