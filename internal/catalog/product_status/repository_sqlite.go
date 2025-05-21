package productstatus

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
	cache []ProductStatus
}

func NewSQLiteRepository(store db.Store) *SQLiteRepository {
	return &SQLiteRepository{store: store}
}

func (repo *SQLiteRepository) Store() db.Store {
	return repo.store
}

func (repo *SQLiteRepository) selectBase() *sqlbuilder.SelectBuilder {
	sb := sqlbuilder.NewSelectBuilder()
	sb.From("product_status as ps")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select status_id, count(*) as count from product group by status_id) as p", "p.status_id = ps.id")
	sb.OrderBy("ps.position", "ps.id DESC")
	return sb
}

func (repo *SQLiteRepository) selectColumns(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
	return sb.Select(
		"ps.id",
		"ps.name",
		"ps.color",
		"coalesce(p.count, 0) as related_products",
	)
}

func (repo *SQLiteRepository) FindMany(ctx context.Context, r w2.GridDataRequest) ([]ProductStatus, int, error) {
	const op = "productstatus.SQLiteRepository.FindMany"

	var total int
	var records []ProductStatus

	sb := repo.selectBase()
	sb.Select("count(*)")
	w2sqlbuilder.Where(sb, r, map[string]string{
		"name": "ps.name",
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
		var record ProductStatus
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) FindAll(ctx context.Context) ([]ProductStatus, error) {
	const op = "productstatus.SQLiteRepository.FindAll"

	if len(repo.cache) > 0 {
		return repo.cache, nil
	}

	var records []ProductStatus

	sb := repo.selectBase()
	repo.selectColumns(sb)

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record ProductStatus
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	repo.cache = records
	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetDropdown(ctx context.Context, r w2.DropdownRequest) ([]ProductStatusDropdown, error) {
	const op = "productstatus.SQLiteRepository.GetDropdown"

	var records []ProductStatusDropdown

	sb := sqlbuilder.Select("id", "name as text", "color")
	sb.From("product_status")
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
		var record ProductStatusDropdown
		if err := rows.Scan(&record.ID, &record.Text, &record.Color); err != nil {
			return nil, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) UpsertMany(ctx context.Context, changes []ProductStatus) error {
	const op = "productstatus.SQLiteRepository.UpsertMany"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		var query string
		var args []any

		if dto.ID == 0 {
			ib := sqlbuilder.InsertInto("product_status")
			ib.Cols("name", "color")
			ib.Values(dto.Name, dto.Color)
			query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
		} else {
			ub := sqlbuilder.Update("product_status")
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

	repo.cache = nil // invalidate cache
	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "productstatus.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("product_status")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	repo.cache = nil // invalidate cache
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) UpdatePositions(ctx context.Context, orderedIDs []int) error {
	const op = "productstatus.SQLiteRepository.UpdatePositions"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for i, id := range orderedIDs {
		ub := sqlbuilder.Update("product_status")
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
