package orderstatus

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
}

func NewSQLiteRepository(store db.Store) *SQLiteRepository {
	return &SQLiteRepository{store: store}
}

func (repo *SQLiteRepository) Store() db.Store {
	return repo.store
}

func (repo *SQLiteRepository) FindMany(ctx context.Context, r w2.GridDataRequest) ([]OrderStatus, int, error) {
	const op = "orderstatus.SQLiteRepository.FindMany"

	var total int
	var records []OrderStatus

	sb := sqlbuilder.Select("count(*)")
	sb.From("order_status as os")
	w2sqlbuilder.Where(sb, r, map[string]string{
		"name": "os.name",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
		"os.id",
		"os.name",
		"os.color",
		"coalesce(o.count, 0) as related_orders",
	)
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select order_status_id, count(*) as count from order_header group by order_status_id) as o", "o.order_status_id = os.id")
	sb.OrderBy("os.position", "os.id DESC")

	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record OrderStatus
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "orderstatus.SQLiteRepository.GetDropdown"

	var records []w2.DropdownValue

	sb := sqlbuilder.Select("id", "name as text")
	sb.From("order_status")
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

func (repo *SQLiteRepository) UpsertMany(ctx context.Context, dtos []OrderStatus) error {
	const op = "orderstatus.SQLiteRepository.UpsertMany"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range dtos {
		var query string
		var args []interface{}

		if dto.ID == 0 {
			ib := sqlbuilder.InsertInto("order_status")
			ib.Cols("name", "color")
			ib.Values(dto.Name, dto.Color)
			query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
		} else {
			ub := sqlbuilder.Update("order_status")
			ub.Where(ub.EQ("id", dto.ID))
			ub.SetMore("updated_at = unixepoch()")
			w2sqlbuilder.SetEditable(ub, dto.Name, "name")
			w2sqlbuilder.SetEditable(ub, dto.Color, "color")
			query, args = ub.BuildWithFlavor(sqlbuilder.SQLite)
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "orderstatus.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("order_status")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) UpdatePositions(ctx context.Context, orderedIDs []int) error {
	const op = "orderstatus.SQLiteRepository.UpdatePositions"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for i, id := range orderedIDs {
		ub := sqlbuilder.Update("order_status")
		ub.Where(ub.EQ("id", id))
		ub.SetMore("updated_at = unixepoch()")
		ub.SetMore(ub.EQ("position", i))

		query, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}
