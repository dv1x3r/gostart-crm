package order

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

func (repo *SQLiteRepository) FindMany(ctx context.Context, r w2.GridDataRequest) ([]Order, int, error) {
	const op = "order.SQLiteRepository.FindMany"

	var total int
	var records []Order

	sb := sqlbuilder.Select("count(*)")
	sb.From("order_header as o")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "order_status as os", "os.id = o.order_status_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "payment_method as pm", "pm.id = o.payment_method_id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, `(
		select
			order_id,
			sum(price * quantity) as total
		from order_line
		group by order_id
	) as l`, "l.order_id = o.id")

	w2sqlbuilder.Where(sb, r, map[string]string{
		"id":           "o.id",
		"full_name":    "concat(o.first_name, ' ', o.last_name)",
		"email":        "o.email",
		"phone_number": "o.phone_number",
		"created_at":   "o.created_at",
		"updated_at":   "o.updated_at",
		"status":       "o.order_status_id",
		"total":        "l.total",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
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
		"os.name as order_status_text",
		"os.color as order_status_color",
		"o.payment_method_id",
		"pm.name as payment_method_text",
		"coalesce(l.total, 0) as total",
	)

	w2sqlbuilder.OrderBy(sb, r, map[string]string{
		"id":           "o.id",
		"full_name":    "concat(o.first_name, ' ', o.last_name)",
		"email":        "o.email",
		"phone_number": "o.phone_number",
		"created_at":   "o.created_at",
		"updated_at":   "o.updated_at",
		"status":       "o.order_status_id",
		"total":        "l.total",
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
		var record Order
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		if record.Lines, err = repo.findLinesByOrderID(ctx, record.ID); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) findLinesByOrderID(ctx context.Context, orderID int) ([]OrderLine, error) {
	const op = "order.SQLiteRepository.findLinesByOrderID"

	var records []OrderLine

	sb := sqlbuilder.Select(
		"id",
		"product_code as code",
		"product_snapshot as product",
		"quantity",
		"price",
		"price * quantity as total",
	).From("order_line")
	sb.OrderBy("price desc, quantity desc")
	sb.Where(sb.EQ("order_id", orderID))

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record OrderLine
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) UpdateByID(ctx context.Context, id int, dto Order) error {
	const op = "order.SQLiteRepository.UpdateByID"

	ub := sqlbuilder.Update("order_header")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")

	w2sqlbuilder.SetEditable(ub, dto.Email, "email")
	w2sqlbuilder.SetEditable(ub, dto.FirstName, "first_name")
	w2sqlbuilder.SetEditable(ub, dto.LastName, "last_name")
	w2sqlbuilder.SetEditable(ub, dto.PhoneNumber, "phone_number")
	w2sqlbuilder.SetEditable(ub, dto.Status.ID, "order_status_id")
	w2sqlbuilder.SetEditable(ub, dto.PaymentMethod.ID, "payment_method_id")

	query, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "order.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("order_header")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}
