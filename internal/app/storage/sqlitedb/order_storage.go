package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Order struct {
	db *sqlx.DB
}

func NewOrder(db *sqlx.DB) *Order {
	return &Order{db: db}
}

func (st *Order) getQueryGetCounter() (string, []any) {
	sb := sqlbuilder.Select("coalesce(count(*), 0) as count")
	sb.From("order_header as o")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "order_status as os", "os.id = o.order_status_id")
	sb.Where("os.in_counter = 1")
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Order) GetCounter(ctx context.Context) (int64, error) {
	const op = "sqlitedb.Order.GetCounter"
	query, args := st.getQueryGetCounter()
	count, _, err := runGet[int64](ctx, st.db, query, args)
	return count, utils.WrapIfErr(op, err)
}

func (st *Order) getQueryFindMany(q storage.FindManyParams) (string, []any) {
	sb := sqlbuilder.Select(
		"o.id",
		"o.email",
		"concat(o.first_name, ' ', o.last_name) as full_name",
		"o.first_name",
		"o.last_name",
		"o.phone_number",
		"o.delivery_address",
		"o.comment",
		"datetime(o.created_at, 'unixepoch', 'localtime') as created_at",
		"datetime(o.updated_at, 'unixepoch', 'localtime') as updated_at",
		"o.order_status_id",
		"os.name as order_status_name",
		"os.color as order_status_color",
		"o.payment_method_id",
		"pm.name as payment_method_name",
		"coalesce(l.total, 0) as total",
		"count(*) over () as count",
	).From("order_header as o")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "order_status as os", "os.id = o.order_status_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "payment_method as pm", "pm.id = o.payment_method_id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, `(
		select order_id, sum(price * quantity) as total
		from order_line
		group by order_id
	) as l`, "l.order_id = o.id")

	allowed := map[string]string{
		"id":           "o.id",
		"full_name":    "concat(o.first_name, ' ', o.last_name)",
		"email":        "o.email",
		"phone_number": "o.phone_number",
		"created_at":   "o.created_at",
		"updated_at":   "o.updated_at",
		"status":       "o.order_status_id",
		"total":        "l.total",
	}

	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowed)
	storage.ApplySorters(sb, q.Sorters, allowed)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Order) FindMany(ctx context.Context, q storage.FindManyParams) ([]model.OrderHeader, int64, error) {
	const op = "sqlitedb.Order.FindMany"

	type LocalResult struct {
		model.OrderHeader
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindMany(q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.OrderHeader, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.OrderHeader, row.Count
	}

	return dto, count, nil
}

func (st *Order) getQueryFindManyLinesByOrderID(orderID int64, q storage.FindManyParams) (string, []any) {
	sb := sqlbuilder.Select(
		"id",
		"product_code as code",
		"product_snapshot as product",
		"quantity",
		"price as price",
		"price * quantity as total",
		"sum(price * quantity) over () as summary",
		"count(*) over () as count",
	).From("order_line")
	sb.OrderBy("price_private desc, quantity desc")
	sb.Where(sb.EQ("order_id", orderID))
	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Order) FindManyLinesByOrderID(ctx context.Context, orderID int64, q storage.FindManyParams) ([]model.OrderLine, int64, model.OrderLineSummary, error) {
	const op = "sqlitedb.Order.FindManyLinesByOrderID"

	summary := model.OrderLineSummary{}

	type LocalResult struct {
		model.OrderLine
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindManyLinesByOrderID(orderID, q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, summary, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.OrderLine, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.OrderLine, row.Count
		summary.Total = row.Summary
	}

	return dto, count, summary, nil
}

func (st *Order) getQueryUpdateByID(dto model.OrderHeader) (string, []any) {
	ub := sqlbuilder.Update("order_header")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Email", "email", dto.Email)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "FirstName", "first_name", dto.FirstName)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "LastName", "last_name", dto.LastName)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "PhoneNumber", "phone_number", dto.PhoneNumber)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "OrderStatusEmbed.StatusID", "order_status_id", dto.StatusID)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "PaymentMethodEmbed.PaymentMethodID", "payment_method_id", dto.PaymentMethodID)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Order) UpdateByID(ctx context.Context, dto model.OrderHeader) error {
	const op = "sqlitedb.Order.UpdateByID"
	query, args := st.getQueryUpdateByID(dto)
	if _, err := runExecAffected(ctx, st.db, query, args); err != nil {
		return utils.WrapIfErr(op, err)
	}
	return nil
}

func (st *Order) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("order_header")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Order) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.Order.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}
