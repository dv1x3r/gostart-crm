package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type AttributeSet struct {
	db *sqlx.DB
}

func NewAttributeSet(db *sqlx.DB) *AttributeSet {
	return &AttributeSet{db: db}
}

func (st *AttributeSet) getQueryFindManyByGroupID(groupID int64, q storage.FindManyParams) (string, []any) {
	sb := sqlbuilder.Select(
		"id",
		"attribute_group_id",
		"name",
		"in_box",
		"in_filter",
		"count(*) over () as count",
	).From("attribute_set")
	sb.Where(sb.EQ("attribute_group_id", groupID))
	sb.OrderBy("position", "id DESC")

	allowed := map[string]string{
		"name": "name",
	}

	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)
	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeSet) FindManyByGroupID(ctx context.Context, groupID int64, q storage.FindManyParams) ([]model.AttributeSet, int64, error) {
	const op = "sqlitedb.AttributeSet.FindManyByGroupID"

	type LocalResult struct {
		model.AttributeSet
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindManyByGroupID(groupID, q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.AttributeSet, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.AttributeSet, row.Count
	}

	return dto, count, nil
}

func (st *AttributeSet) getQueryGetByID(id int64) (string, []any) {
	sb := sqlbuilder.Select(
		"id",
		"attribute_group_id",
		"name",
		"in_box",
		"in_filter",
	).From("attribute_set")
	sb.Where(sb.EQ("id", id))
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeSet) GetByID(ctx context.Context, id int64) (model.AttributeSet, bool, error) {
	const op = "sqlitedb.AttributeSet.GetByID"
	query, args := st.getQueryGetByID(id)
	row, ok, err := runGet[model.AttributeSet](ctx, st.db, query, args)
	return row, ok, utils.WrapIfErr(op, err)
}

func (st *AttributeSet) getQueryInsert(groupID int64, dto model.AttributeSet) (string, []any) {
	ib := sqlbuilder.InsertInto("attribute_set")
	ib.Cols("attribute_group_id", "name", "in_box", "in_filter")
	ib.Values(groupID, dto.Name, dto.InBox, dto.InFilter)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeSet) getQueryUpdateByID(dto model.AttributeSet) (string, []any) {
	ub := sqlbuilder.Update("attribute_set")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "InBox", "in_box", dto.InBox)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "InFilter", "in_filter", dto.InFilter)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeSet) UpsertMany(ctx context.Context, groupID int64, dtos []model.AttributeSet) (int64, error) {
	const op = "sqlitedb.AttributeSet.UpsertMany"

	queries, args := make([]string, len(dtos)), make([][]any, len(dtos))
	for i, dto := range dtos {
		if dto.ID == 0 {
			queries[i], args[i] = st.getQueryInsert(groupID, dto)
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

func (st *AttributeSet) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("attribute_set")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeSet) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.AttributeSet.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *AttributeSet) getQueryUpdatePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("attribute_set")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeSet) UpdatePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.AttributeSet.UpdatePositions"

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
