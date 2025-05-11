package paymentmethod

import (
	"context"
	"database/sql"

	"gostart-crm/internal/lib/db"
	"gostart-crm/internal/lib/wrap"

	"github.com/dv1x3r/w2go/w2"
	"github.com/dv1x3r/w2go/w2sql/w2sqlbuilder"
	"github.com/huandu/go-sqlbuilder"
)

type SQLiteRepository struct {
	store db.Store
	cache []PaymentMethod
}

func NewSQLiteRepository(store db.Store) *SQLiteRepository {
	return &SQLiteRepository{store: store}
}

func (repo *SQLiteRepository) Store() db.Store {
	return repo.store
}

func (repo *SQLiteRepository) selectBase() *sqlbuilder.SelectBuilder {
	sb := sqlbuilder.NewSelectBuilder()
	sb.From("payment_method as pm")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select payment_method_id, count(*) as count from order_header group by payment_method_id) as o", "o.payment_method_id = pm.id")
	sb.OrderBy("pm.position", "pm.id DESC")
	return sb
}

func (repo *SQLiteRepository) selectColumns(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
	return sb.Select(
		"pm.id",
		"pm.name",
		"coalesce(o.count, 0) as related_orders",
	)
}

func (repo *SQLiteRepository) FindMany(ctx context.Context, r w2.GridDataRequest) ([]PaymentMethod, int, error) {
	const op = "paymentmethod.SQLiteRepository.FindMany"

	var total int
	var records []PaymentMethod

	sb := repo.selectBase()
	sb.Select("count(*)")
	w2sqlbuilder.Where(sb, r, map[string]string{
		"name": "pm.name",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	repo.selectColumns(sb)
	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record PaymentMethod
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) FindAll(ctx context.Context) ([]PaymentMethod, error) {
	const op = "paymentmethod.SQLiteRepository.FindAll"

	if len(repo.cache) > 0 {
		return repo.cache, nil
	}

	var records []PaymentMethod

	sb := repo.selectBase()
	repo.selectColumns(sb)

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record PaymentMethod
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	repo.cache = records
	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "paymentmethod.SQLiteRepository.GetDropdown"

	var records []w2.DropdownValue

	sb := sqlbuilder.Select("id", "name as text")
	sb.From("payment_method")
	if r.Search != "" {
		sb.Where(sb.Like("name", "%"+r.Search+"%"))
	}
	sb.OrderBy("position")
	sb.Limit(r.Max)

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record w2.DropdownValue
		if err := rows.Scan(&record.ID, &record.Text); err != nil {
			return nil, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) UpsertMany(ctx context.Context, changes []PaymentMethod) error {
	const op = "paymentmethod.SQLiteRepository.UpsertMany"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		var query string
		var args []any

		if dto.ID == 0 {
			ib := sqlbuilder.InsertInto("payment_method")
			ib.Cols("name")
			ib.Values(dto.Name)
			query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
		} else {
			ub := sqlbuilder.Update("payment_method")
			ub.Where(ub.EQ("id", dto.ID))
			ub.SetMore("updated_at = unixepoch()")
			w2sqlbuilder.SetEditable(ub, dto.Name, "name")
			query, args = ub.BuildWithFlavor(sqlbuilder.SQLite)
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	repo.cache = nil // invalidate cache
	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "paymentmethod.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("payment_method")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	repo.cache = nil // invalidate cache
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) UpdatePositions(ctx context.Context, orderedIDs []int) error {
	const op = "paymentmethod.SQLiteRepository.UpdatePositions"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for i, id := range orderedIDs {
		ub := sqlbuilder.Update("payment_method")
		ub.Where(ub.EQ("id", id))
		ub.SetMore("updated_at = unixepoch()")
		ub.SetMore(ub.EQ("position", i))

		query, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	repo.cache = nil // invalidate cache
	return wrap.IfErr(op, tx.Commit())
}
