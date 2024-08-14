package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Attribute struct {
	db *sqlx.DB
}

func NewAttribute(db *sqlx.DB) *Attribute {
	return &Attribute{db: db}
}

func (st *Attribute) getQueryFindManyGroups(q storage.FindManyParams) (string, []any) {
	sb := sqlbuilder.Select(
		"id",
		"name",
		"count(*) over () as count",
	)
	sb.From("attribute_group")
	sb.OrderBy("name")

	allowed := map[string]string{
		"name": "name",
	}

	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)
	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) FindManyGroups(ctx context.Context, q storage.FindManyParams) ([]model.AttributeGroup, int64, error) {
	const op = "sqlitedb.Attribute.FindManyGroups"

	type LocalResult struct {
		model.AttributeGroup
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindManyGroups(q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.AttributeGroup, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.AttributeGroup, row.Count
	}

	return dto, count, utils.WrapIfErr(op, err)
}

func (st *Attribute) getQueryGetGroupDropdown(search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text")
	sb.From("attribute_group")
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("name")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) GetGroupDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.Attribute.GetGroupDropdown"
	query, args := st.getQueryGetGroupDropdown(search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *Attribute) getQueryInsertGroup(dto model.AttributeGroup) (string, []any) {
	ib := sqlbuilder.InsertInto("attribute_group")
	ib.Cols("name")
	ib.Values(dto.Name)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) getQueryUpdateGroupByID(dto model.AttributeGroup) (string, []any) {
	ub := sqlbuilder.Update("attribute_group")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) UpsertGroup(ctx context.Context, dto model.AttributeGroup) error {
	const op = "sqlitedb.Attribute.UpsertGroup"
	if dto.ID == 0 {
		query, args := st.getQueryInsertGroup(dto)
		_, err := runExecAffected(ctx, st.db, query, args)
		return utils.WrapIfErr(op, err)
	} else {
		query, args := st.getQueryUpdateGroupByID(dto)
		_, err := runExecAffected(ctx, st.db, query, args)
		return utils.WrapIfErr(op, err)
	}
}

func (st *Attribute) getQueryDeleteGroupByID(id int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("attribute_group")
	dlb.Where(dlb.EQ("id", id))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) DeleteGroupByID(ctx context.Context, id int64) error {
	const op = "sqlitedb.Attribute.DeleteGroupByID"
	query, args := st.getQueryDeleteGroupByID(id)
	_, err := runExecAffected(ctx, st.db, query, args)
	return utils.WrapIfErr(op, err)
}

func (st *Attribute) getQueryFindManySetsByGroupID(groupID int64, q storage.FindManyParams) (string, []any) {
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

func (st *Attribute) FindManySetsByGroupID(ctx context.Context, groupID int64, q storage.FindManyParams) ([]model.AttributeSet, int64, error) {
	const op = "sqlitedb.Attribute.FindManySets"

	type LocalResult struct {
		model.AttributeSet
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindManySetsByGroupID(groupID, q)
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

func (st *Attribute) getQueryGetSetByID(id int64) (string, []any) {
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

func (st *Attribute) GetSetByID(ctx context.Context, id int64) (model.AttributeSet, bool, error) {
	const op = "sqlitedb.Attribute.GetSetByID"
	query, args := st.getQueryGetSetByID(id)
	row, ok, err := runGet[model.AttributeSet](ctx, st.db, query, args)
	return row, ok, utils.WrapIfErr(op, err)
}

func (st *Attribute) getQueryInsertSet(groupID int64, dto model.AttributeSet) (string, []any) {
	ib := sqlbuilder.InsertInto("attribute_set")
	ib.Cols("attribute_group_id", "name", "in_box", "in_filter")
	ib.Values(groupID, dto.Name, dto.InBox, dto.InFilter)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) getQueryUpdateSetByID(dto model.AttributeSet) (string, []any) {
	ub := sqlbuilder.Update("attribute_set")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "InBox", "in_box", dto.InBox)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "InFilter", "in_filter", dto.InFilter)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) UpsertManySets(ctx context.Context, groupID int64, dtos []model.AttributeSet) (int64, error) {
	const op = "sqlitedb.Attribute.UpsertManySets"

	queries, args := make([]string, len(dtos)), make([][]any, len(dtos))
	for i, dto := range dtos {
		if dto.ID == 0 {
			queries[i], args[i] = st.getQueryInsertSet(groupID, dto)
		} else {
			queries[i], args[i] = st.getQueryUpdateSetByID(dto)
		}
	}

	affected, err := runExecAffectedRangeNewTx(ctx, st.db, queries, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	return affected, nil
}

func (st *Attribute) getQueryDeleteManySetsByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("attribute_set")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) DeleteManySetsByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.Attribute.DeleteManySetsByID"
	query, args := st.getQueryDeleteManySetsByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *Attribute) getQueryUpdateSetPosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("attribute_set")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) UpdateSetPositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.Attribute.UpdateSetPositions"

	queries, args := make([]string, len(orderedIDs)), make([][]any, len(orderedIDs))
	for i, id := range orderedIDs {
		queries[i], args[i] = st.getQueryUpdateSetPosition(id, int64(i))
	}

	affected, err := runExecAffectedRangeNewTx(ctx, st.db, queries, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	return affected, nil
}

func (st *Attribute) getQueryFindManyValuesBySetID(setID int64, q storage.FindManyParams) (string, []any) {
	sb := sqlbuilder.Select(
		"av.id",
		"av.attribute_set_id",
		"av.name",
		"coalesce(p.count, 0) as related_products",
		"count(*) over () as count",
	).From("attribute_value as av")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select attribute_value_id, count(*) as count from product_attribute group by attribute_value_id) as p", "p.attribute_value_id = av.id")
	sb.Where(sb.EQ("av.attribute_set_id", setID))
	sb.OrderBy("av.position", "av.id DESC")

	allowed := map[string]string{
		"name": "av.name",
	}

	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)
	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) FindManyValuesBySetID(ctx context.Context, setID int64, q storage.FindManyParams) ([]model.AttributeValue, int64, error) {
	const op = "sqlitedb.Attribute.FindManyValues"

	type LocalResult struct {
		model.AttributeValue
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindManyValuesBySetID(setID, q)
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

func (st *Attribute) getQueryGetValueDropdownBySetID(setID int64, search string, max int) (string, []any) {
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

func (st *Attribute) GetValueDropdownBySetID(ctx context.Context, setID int64, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.Attribute.GetValueDropdownBySetID"
	query, args := st.getQueryGetValueDropdownBySetID(setID, search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *Attribute) getQueryGetValueByID(id int64) (string, []any) {
	sb := sqlbuilder.Select("id", "attribute_set_id", "name")
	sb.From("attribute_value")
	sb.Where(sb.EQ("id", id))
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) GetValueByID(ctx context.Context, id int64) (model.AttributeValue, bool, error) {
	const op = "sqlitedb.Media.GetValueByID"
	query, args := st.getQueryGetValueByID(id)
	row, ok, err := runGet[model.AttributeValue](ctx, st.db, query, args)
	return row, ok, err
}

func (st *Attribute) getQueryInsertValue(setID int64, dto model.AttributeValue) (string, []any) {
	ib := sqlbuilder.InsertInto("attribute_value")
	ib.Cols("attribute_set_id", "name")
	ib.Values(setID, dto.Name)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) getQueryUpdateValueByID(dto model.AttributeValue) (string, []any) {
	ub := sqlbuilder.Update("attribute_value")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) UpsertManyValues(ctx context.Context, setID int64, dtos []model.AttributeValue) (int64, error) {
	const op = "sqlitedb.Attribute.UpsertManyValues"

	queries, args := make([]string, len(dtos)), make([][]any, len(dtos))
	for i, dto := range dtos {
		if dto.ID == 0 {
			queries[i], args[i] = st.getQueryInsertValue(setID, dto)
		} else {
			queries[i], args[i] = st.getQueryUpdateValueByID(dto)
		}
	}

	affected, err := runExecAffectedRangeNewTx(ctx, st.db, queries, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	return affected, nil
}

func (st *Attribute) getQueryDeleteManyValuesByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("attribute_value")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) DeleteManyValuesByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.Attribute.DeleteManyValuesByID"
	query, args := st.getQueryDeleteManyValuesByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *Attribute) getQueryUpdateValuePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("attribute_value")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Attribute) UpdateValuePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.Attribute.UpdateValuePositions"

	queries, args := make([]string, len(orderedIDs)), make([][]any, len(orderedIDs))
	for i, id := range orderedIDs {
		queries[i], args[i] = st.getQueryUpdateValuePosition(id, int64(i))
	}

	affected, err := runExecAffectedRangeNewTx(ctx, st.db, queries, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	return affected, nil
}
