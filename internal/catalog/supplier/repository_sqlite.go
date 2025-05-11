package supplier

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

func (repo *SQLiteRepository) FindMany(ctx context.Context, r w2.GridDataRequest) ([]Supplier, int, error) {
	const op = "supplier.SQLiteRepository.FindMany"

	var total int
	var records []Supplier

	sb := sqlbuilder.Select("count(*)")
	sb.From("supplier as s")

	sb.JoinWithOption(sqlbuilder.LeftJoin, `(
		select
			p.supplier_id,
			count(*) as related_products,
			sum(iif(p.quantity > 0 and p.is_published = 1 and c.is_published = 1 and s.is_published = 1, 1, 0)) as published_products
		from product as p
		join category as c on c.id = p.category_id
		join supplier as s on s.id = p.supplier_id
		group by supplier_id
	) as p`, "p.supplier_id = s.id")

	w2sqlbuilder.Where(sb, r, map[string]string{
		"code":               "s.code",
		"name":               "s.name",
		"description":        "s.description",
		"is_published":       "s.is_published",
		"related_products":   "p.related_products",
		"published_products": "p.published_products",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
		"s.id",
		"s.code",
		"s.name",
		"s.description",
		"s.is_published",
		"coalesce(p.related_products, 0) as related_products",
		"coalesce(p.published_products, 0) as published_products",
	)

	sb.OrderBy("position", "id DESC")

	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Supplier
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "supplier.SQLiteRepository.GetDropdown"

	var records []w2.DropdownValue

	sb := sqlbuilder.Select("id", "name as text")
	sb.From("supplier")
	if r.Search != "" {
		sb.Where(sb.Like("name_lv", "%"+r.Search+"%"))
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

func (repo *SQLiteRepository) UpsertMany(ctx context.Context, changes []Supplier) error {
	const op = "supplier.SQLiteRepository.UpsertMany"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		var query string
		var args []any

		if dto.ID == 0 {
			ib := sqlbuilder.InsertInto("supplier")
			ib.Cols("slug", "code", "name", "description", "is_published")
			ib.Values(dto.Slugify(), dto.Code, dto.Name, dto.Description, dto.IsPublished.V)
			query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
		} else {
			ub := sqlbuilder.Update("supplier")
			ub.Where(ub.EQ("id", dto.ID))
			ub.SetMore("updated_at = unixepoch()")

			slug := dto.Slugify()
			if slug != "" {
				ub.SetMore(ub.EQ("slug", slug))
			}

			w2sqlbuilder.SetEditable(ub, dto.Code, "code")
			w2sqlbuilder.SetEditable(ub, dto.Name, "name")
			w2sqlbuilder.SetEditable(ub, dto.Description, "description")
			w2sqlbuilder.SetEditable(ub, dto.IsPublished, "is_published")

			query, args = ub.BuildWithFlavor(sqlbuilder.SQLite)
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "supplier.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("supplier")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) UpdatePositions(ctx context.Context, orderedIDs []int) error {
	const op = "sqlitedb.Supplier.UpdatePositions"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for i, id := range orderedIDs {
		ub := sqlbuilder.Update("supplier")
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
