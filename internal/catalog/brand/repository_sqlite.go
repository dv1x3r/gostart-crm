package brand

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

func (repo *SQLiteRepository) FindMany(ctx context.Context, r w2.GridDataRequest) ([]Brand, int, error) {
	const op = "brand.SQLiteRepository.FindMany"

	var total int
	var records []Brand

	sb := sqlbuilder.Select("count(*)")
	sb.From("brand as b")

	sb.JoinWithOption(sqlbuilder.LeftJoin, `(
		select
			p.brand_id,
			count(*) as related_products,
			sum(iif(p.quantity > 0 and p.is_published = 1 and c.is_published = 1 and s.is_published = 1, 1, 0)) as published_products
		from product as p
		join category as c on c.id = p.category_id
		join supplier as s on s.id = p.supplier_id
		group by brand_id
	) as p`, "p.brand_id = b.id")

	w2sqlbuilder.Where(sb, r, map[string]string{
		"name":               "b.name",
		"related_products":   "p.related_products",
		"published_products": "p.published_products",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
		"b.id",
		"b.name",
		"coalesce(p.related_products, 0) as related_products",
		"coalesce(p.published_products, 0) as published_products",
	)

	w2sqlbuilder.OrderBy(sb, r, map[string]string{
		"id":                 "b.id",
		"name":               "b.name",
		"related_products":   "p.related_products",
		"published_products": "p.published_products",
	})

	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Brand
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "brand.SQLiteRepository.GetDropdown"

	var records []w2.DropdownValue

	sb := sqlbuilder.Select("id", "name as text")
	sb.From("brand")
	if r.Search != "" {
		sb.Where(sb.Like("name", "%"+r.Search+"%"))
	}
	sb.OrderBy("name")
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

func (repo *SQLiteRepository) UpsertMany(ctx context.Context, changes []Brand) error {
	const op = "brand.SQLiteRepository.UpsertMany"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		var query string
		var args []any

		if dto.ID == 0 {
			ib := sqlbuilder.InsertInto("brand")
			ib.Cols("name")
			ib.Values(dto.Name)
			query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
		} else {
			ub := sqlbuilder.Update("brand")
			ub.Where(ub.EQ("id", dto.ID))
			ub.SetMore("updated_at = unixepoch()")
			w2sqlbuilder.SetEditable(ub, dto.Name, "name")
			query, args = ub.BuildWithFlavor(sqlbuilder.SQLite)
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "brand.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("brand")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}
