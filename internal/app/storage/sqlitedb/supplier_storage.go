package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Supplier struct {
	db *sqlx.DB
}

func NewSupplier(db *sqlx.DB) *Supplier {
	return &Supplier{db: db}
}

func (st *Supplier) getQuerySelectBase(extra ...string) *sqlbuilder.SelectBuilder {
	columns := []string{
		"s.id",
		"s.slug",
		"s.code",
		"s.name",
		"s.description",
		"s.is_published",
		"coalesce(p.count, 0) as related_products",
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(append(columns, extra...)...).From("supplier as s")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "(select supplier_id, count(*) as count from product group by supplier_id) as p", "p.supplier_id = s.id")
	sb.OrderBy("position", "id DESC")
	return sb
}

func (st *Supplier) getQueryFindMany(q storage.FindManyParams) (string, []any) {
	sb := st.getQuerySelectBase("count(*) over () as count")

	allowedFilters := map[string]string{
		"code":             "s.code",
		"name":             "s.name",
		"description":      "s.description",
		"is_published":     "s.is_published",
		"related_products": "coalesce(p.count, 0)",
	}

	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowedFilters)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Supplier) FindMany(ctx context.Context, q storage.FindManyParams) ([]model.Supplier, int64, error) {
	const op = "sqlitedb.Supplier.FindMany"

	type LocalResult struct {
		model.Supplier
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindMany(q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.Supplier, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.Supplier, row.Count
	}

	return dto, count, nil
}

func (st *Supplier) getQueryGetDropdown(search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text")
	sb.From("supplier")
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("position")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Supplier) GetDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.Supplier.GetDropdown"
	query, args := st.getQueryGetDropdown(search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *Supplier) getQueryGetByID(id int64) (string, []any) {
	sb := st.getQuerySelectBase()
	sb.Where(sb.EQ("id", id))
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Supplier) GetByID(ctx context.Context, id int64) (model.Supplier, bool, error) {
	const op = "sqlitedb.Supplier.GetByID"
	query, args := st.getQueryGetByID(id)
	row, ok, err := runGet[model.Supplier](ctx, st.db, query, args)
	return row, ok, utils.WrapIfErr(op, err)
}

func (st *Supplier) getQueryInsert(dto model.Supplier) (string, []any) {
	ib := sqlbuilder.InsertInto("supplier")
	ib.Cols("slug", "code", "name", "description", "is_published")
	ib.Values(dto.Slugify(), dto.Code, dto.Name, dto.Description, dto.IsPublished)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Supplier) getQueryUpdateByID(dto model.Supplier) (string, []any) {
	ub := sqlbuilder.Update("supplier")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	slug := dto.Slugify()
	if slug != "" {
		ub.SetMore(ub.EQ("slug", slug))
	}
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Code", "code", dto.Code)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Description", "description", dto.Description)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "IsPublished", "is_published", dto.IsPublished)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Supplier) UpsertMany(ctx context.Context, dtos []model.Supplier) (int64, error) {
	const op = "sqlitedb.Supplier.UpsertMany"

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

	return affected, utils.WrapIfErr(op, err)
}

func (st *Supplier) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("supplier")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Supplier) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.Supplier.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *Supplier) getQueryUpdatePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("supplier")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Supplier) UpdatePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.Supplier.UpdatePositions"

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
