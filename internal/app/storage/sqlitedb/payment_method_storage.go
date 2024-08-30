package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type PaymentMethod struct {
	db *sqlx.DB
}

func NewPaymentMethod(db *sqlx.DB) *PaymentMethod {
	return &PaymentMethod{db: db}
}

func (st *PaymentMethod) getQuerySelectBase(extra ...string) *sqlbuilder.SelectBuilder {
	columns := []string{
		"pm.id",
		"pm.name",
		"coalesce(o.count, 0) as related_orders",
	}

	sb := sqlbuilder.Select(append(columns, extra...)...).From("payment_method as pm")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select payment_method_id, count(*) as count from order_header group by payment_method_id) as o", "o.payment_method_id = pm.id")
	sb.OrderBy("pm.position", "pm.id DESC")
	return sb
}

func (st *PaymentMethod) getQueryFindMany(q storage.FindManyParams) (string, []any) {
	sb := st.getQuerySelectBase("count(*) over () as count")

	allowed := map[string]string{
		"name": "pm.name",
	}

	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *PaymentMethod) FindMany(ctx context.Context, q storage.FindManyParams) ([]model.PaymentMethod, int64, error) {
	const op = "sqlitedb.PaymentMethod.FindMany"

	type LocalResult struct {
		model.PaymentMethod
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindMany(q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.PaymentMethod, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.PaymentMethod, row.Count
	}

	return dto, count, nil
}

func (st *PaymentMethod) getQueryGetDropdown(search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text")
	sb.From("payment_method")
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("position")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *PaymentMethod) GetDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.PaymentMethod.GetDropdown"
	query, args := st.getQueryGetDropdown(search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *PaymentMethod) getQueryInsert(dto model.PaymentMethod) (string, []any) {
	ib := sqlbuilder.InsertInto("payment_method")
	ib.Cols("name")
	ib.Values(dto.Name)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *PaymentMethod) getQueryUpdateByID(dto model.PaymentMethod) (string, []any) {
	ub := sqlbuilder.Update("payment_method")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *PaymentMethod) UpsertMany(ctx context.Context, dtos []model.PaymentMethod) (int64, error) {
	const op = "sqlitedb.PaymentMethod.UpsertMany"

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

func (st *PaymentMethod) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("payment_method")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *PaymentMethod) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.PaymentMethod.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *PaymentMethod) getQueryUpdatePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("payment_method")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *PaymentMethod) UpdatePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.PaymentMethod.UpdatePositions"

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
