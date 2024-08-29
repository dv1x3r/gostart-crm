package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type AttributeValue struct {
	db *sqlx.DB
}

func NewAttributeValue(db *sqlx.DB) *AttributeValue {
	return &AttributeValue{db: db}
}

func (st *AttributeValue) getQueryFindManyBySetID(setID int64, q storage.FindManyParams) (string, []any) {
	sb := sqlbuilder.Select(
		"av.id",
		"av.attribute_set_id",
		"av.name",
		"coalesce(p.related_products, 0) as related_products",
		"coalesce(p.published_products, 0) as published_products",
		"count(*) over () as count",
	).From("attribute_value as av")
	sb.JoinWithOption(sqlbuilder.LeftJoin, `(
		select
			av.attribute_value_id,
			count(*) as related_products,
			sum(iif(p.quantity > 0 and p.is_published = 1 and c.is_published = 1 and s.is_published = 1, 1, 0)) as published_products
		from product_attribute as av
		join product as p on p.id = av.product_id
		join category as c on c.id = p.category_id
		join supplier as s on s.id = p.supplier_id
		group by av.attribute_value_id
	) as p`, "p.attribute_value_id = av.id")
	sb.Where(sb.EQ("av.attribute_set_id", setID))
	sb.OrderBy("av.position", "av.id DESC")

	allowed := map[string]string{
		"name": "av.name",
	}

	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)
	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeValue) FindManyBySetID(ctx context.Context, setID int64, q storage.FindManyParams) ([]model.AttributeValue, int64, error) {
	const op = "sqlitedb.AttributeValue.FindMany"

	type LocalResult struct {
		model.AttributeValue
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindManyBySetID(setID, q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.AttributeValue, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.AttributeValue, row.Count
	}

	return dto, count, nil
}

func (st *AttributeValue) getQueryGetDropdownBySetID(setID int64, search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text")
	sb.From("attribute_value")
	sb.Where(sb.EQ("attribute_set_id", setID))
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("position", "id DESC")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeValue) GetDropdownBySetID(ctx context.Context, setID int64, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.AttributeValue.GetDropdownBySetID"
	query, args := st.getQueryGetDropdownBySetID(setID, search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *AttributeValue) getQueryGetByID(id int64) (string, []any) {
	sb := sqlbuilder.Select("id", "attribute_set_id", "name")
	sb.From("attribute_value")
	sb.Where(sb.EQ("id", id))
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeValue) GetByID(ctx context.Context, id int64) (model.AttributeValue, bool, error) {
	const op = "sqlitedb.AttributeValue.GetByID"
	query, args := st.getQueryGetByID(id)
	row, ok, err := runGet[model.AttributeValue](ctx, st.db, query, args)
	return row, ok, err
}

func (st *AttributeValue) getQueryInsert(setID int64, dto model.AttributeValue) (string, []any) {
	ib := sqlbuilder.InsertInto("attribute_value")
	ib.Cols("attribute_set_id", "name")
	ib.Values(setID, dto.Name)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeValue) getQueryUpdateByID(dto model.AttributeValue) (string, []any) {
	ub := sqlbuilder.Update("attribute_value")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeValue) UpsertMany(ctx context.Context, setID int64, dtos []model.AttributeValue) (int64, error) {
	const op = "sqlitedb.AttributeValue.UpsertMany"

	queries, args := make([]string, len(dtos)), make([][]any, len(dtos))
	for i, dto := range dtos {
		if dto.ID == 0 {
			queries[i], args[i] = st.getQueryInsert(setID, dto)
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

func (st *AttributeValue) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("attribute_value")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeValue) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.AttributeValue.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *AttributeValue) getQueryUpdatePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("attribute_value")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeValue) UpdatePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.AttributeValue.UpdatePositions"

	queries, args := make([]string, len(orderedIDs)), make([][]any, len(orderedIDs))
	for i, id := range orderedIDs {
		queries[i], args[i] = st.getQueryUpdatePosition(id, int64(i))
	}

	affected, err := runExecAffectedRangeNewTx(ctx, st.db, queries, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	return affected, nil
}
