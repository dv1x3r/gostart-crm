package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type OrderStatus struct {
	db *sqlx.DB
}

func NewOrderStatus(db *sqlx.DB) *OrderStatus {
	return &OrderStatus{db: db}
}

func (st *OrderStatus) getQuerySelectBase(extra ...string) *sqlbuilder.SelectBuilder {
	columns := []string{
		"os.id",
		"os.name",
		"os.color",
		"coalesce(o.count, 0) as related_orders",
	}

	sb := sqlbuilder.Select(append(columns, extra...)...).From("order_status as os")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select order_status_id, count(*) as count from order_header group by order_status_id) as o", "o.order_status_id = os.id")
	sb.OrderBy("os.position", "os.id DESC")
	return sb
}

func (st *OrderStatus) getQueryFindMany(q storage.FindManyParams) (string, []any) {
	sb := st.getQuerySelectBase("count(*) over () as count")

	allowed := map[string]string{
		"name": "os.name",
	}

	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *OrderStatus) FindMany(ctx context.Context, q storage.FindManyParams) ([]model.OrderStatus, int64, error) {
	const op = "sqlitedb.OrderStatus.FindMany"

	type LocalResult struct {
		model.OrderStatus
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindMany(q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.OrderStatus, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.OrderStatus, row.Count
	}

	return dto, count, nil
}

func (st *OrderStatus) FindAll(ctx context.Context) ([]model.OrderStatus, error) {
	const op = "sqlitedb.OrderStatus.FindAll"
	query, args := st.getQuerySelectBase().BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := runSelect[model.OrderStatus](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *OrderStatus) getQueryGetDropdown(search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text")
	sb.From("order_status")
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("position")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *OrderStatus) GetDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.OrderStatus.GetDropdown"
	query, args := st.getQueryGetDropdown(search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *OrderStatus) getQueryInsert(dto model.OrderStatus) (string, []any) {
	ib := sqlbuilder.InsertInto("order_status")
	ib.Cols("name", "color")
	ib.Values(dto.Name, dto.Color)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *OrderStatus) getQueryUpdateByID(dto model.OrderStatus) (string, []any) {
	ub := sqlbuilder.Update("order_status")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Color", "color", dto.Color)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *OrderStatus) UpsertMany(ctx context.Context, dtos []model.OrderStatus) (int64, error) {
	const op = "sqlitedb.OrderStatus.UpsertMany"

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

func (st *OrderStatus) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("order_status")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *OrderStatus) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.OrderStatus.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *OrderStatus) getQueryUpdatePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("order_status")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *OrderStatus) UpdatePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.OrderStatus.UpdatePositions"

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
