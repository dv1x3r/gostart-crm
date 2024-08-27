package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Brand struct {
	db *sqlx.DB
}

func NewBrand(db *sqlx.DB) *Brand {
	return &Brand{db: db}
}

func (st *Brand) getQueryFindMany(q storage.FindManyParams) (string, []any) {
	sb := sqlbuilder.Select(
		"b.id",
		"b.name",
		"coalesce(p.related_products, 0) as related_products",
		"coalesce(p.published_products, 0) as published_products",
		"count(*) over () as count",
	).From("brand as b")
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

	allowed := map[string]string{
		"name":               "b.name",
		"related_products":   "p.related_products",
		"published_products": "p.published_products",
	}

	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)
	storage.ApplySorters(sb, q.Sorters, allowed)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Brand) FindMany(ctx context.Context, q storage.FindManyParams) ([]model.Brand, int64, error) {
	const op = "sqlitedb.Brand.FindMany"

	type LocalResult struct {
		model.Brand
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindMany(q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.Brand, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.Brand, row.Count
	}

	return dto, count, nil
}

func (st *Brand) getQueryGetDropdown(search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text")
	sb.From("brand")
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("name")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Brand) GetDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.Brand.GetDropdown"
	query, args := st.getQueryGetDropdown(search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *Brand) getQueryInsert(dto model.Brand) (string, []any) {
	ib := sqlbuilder.InsertInto("brand")
	ib.Cols("name")
	ib.Values(dto.Name)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Brand) getQueryUpdateByID(dto model.Brand) (string, []any) {
	ub := sqlbuilder.Update("brand")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Brand) UpsertMany(ctx context.Context, dtos []model.Brand) (int64, error) {
	const op = "sqlitedb.Brand.UpsertMany"

	queries, args := make([]string, len(dtos)), make([][]any, len(dtos))
	for i, dto := range dtos {
		if dto.ID == 0 {
			queries[i], args[i] = st.getQueryInsert(dto)
		} else {
			queries[i], args[i] = st.getQueryUpdateByID(dto)
		}
	}

	affected, err := runExecAffectedRangeNewTx(ctx, st.db, queries, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	return affected, nil
}

func (st *Brand) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("brand")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Brand) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.Brand.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}
