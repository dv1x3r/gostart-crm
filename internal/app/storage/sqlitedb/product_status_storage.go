package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type ProductStatus struct {
	db    *sqlx.DB
	cache []model.ProductStatus
}

func NewProductStatus(db *sqlx.DB) *ProductStatus {
	return &ProductStatus{db: db}
}

func (st *ProductStatus) getQuerySelectBase(extra ...string) *sqlbuilder.SelectBuilder {
	columns := []string{
		"ps.id",
		"ps.name",
		"ps.color",
		"coalesce(p.count, 0) as related_products",
	}

	sb := sqlbuilder.Select(append(columns, extra...)...).From("product_status as ps")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select status_id, count(*) as count from product group by status_id) as p", "p.status_id = ps.id")
	sb.OrderBy("ps.position", "ps.id DESC")
	return sb
}

func (st *ProductStatus) getQueryFindMany(q storage.FindManyParams) (string, []any) {
	sb := st.getQuerySelectBase("count(*) over () as count")

	allowed := map[string]string{
		"name": "ps.name",
	}

	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *ProductStatus) FindMany(ctx context.Context, q storage.FindManyParams) ([]model.ProductStatus, int64, error) {
	const op = "sqlitedb.ProductStatus.FindMany"

	type LocalResult struct {
		model.ProductStatus
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindMany(q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.ProductStatus, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.ProductStatus, row.Count
	}

	return dto, count, nil
}

func (st *ProductStatus) FindAll(ctx context.Context) ([]model.ProductStatus, error) {
	const op = "sqlitedb.ProductStatus.FindAll"

	if len(st.cache) > 0 {
		return st.cache, nil
	}

	query, args := st.getQuerySelectBase().BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := runSelect[model.ProductStatus](ctx, st.db, query, args)

	st.cache = rows
	return rows, utils.WrapIfErr(op, err)
}

func (st *ProductStatus) getQueryGetDropdown(search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text", "color")
	sb.From("product_status")
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("position")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *ProductStatus) GetDropdown(ctx context.Context, search string, max int) ([]model.ProductStatusDropdown, error) {
	const op = "sqlitedb.ProductStatus.GetDropdown"
	query, args := st.getQueryGetDropdown(search, max)
	rows, err := runSelect[model.ProductStatusDropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *ProductStatus) getQueryInsert(dto model.ProductStatus) (string, []any) {
	ib := sqlbuilder.InsertInto("product_status")
	ib.Cols("name", "color")
	ib.Values(dto.Name, dto.Color)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *ProductStatus) getQueryUpdateByID(dto model.ProductStatus) (string, []any) {
	ub := sqlbuilder.Update("product_status")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Color", "color", dto.Color)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *ProductStatus) UpsertMany(ctx context.Context, dtos []model.ProductStatus) (int64, error) {
	const op = "sqlitedb.ProductStatus.UpsertMany"

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

	st.cache = st.cache[:0] // invalidate cache
	return affected, nil
}

func (st *ProductStatus) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("product_status")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *ProductStatus) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.ProductStatus.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	st.cache = st.cache[:0] // invalidate cache
	return affected, utils.WrapIfErr(op, err)
}

func (st *ProductStatus) getQueryUpdatePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("product_status")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *ProductStatus) UpdatePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.ProductStatus.UpdatePositions"

	queries, args := make([]string, len(orderedIDs)), make([][]any, len(orderedIDs))
	for i, id := range orderedIDs {
		queries[i], args[i] = st.getQueryUpdatePosition(id, int64(i))
	}

	affected, err := runExecAffectedRangeNewTx(ctx, st.db, queries, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	st.cache = st.cache[:0] // invalidate cache
	return affected, nil
}
