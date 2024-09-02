package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Category struct {
	db *sqlx.DB
}

func NewCategory(db *sqlx.DB) *Category {
	return &Category{db: db}
}

const CategoryCTE = `
	with recursive category_cte as (
		select
			id,
			name,
			slug || '/' as slug,
			attribute_group_id,
			printf('%04d', mp_position) || '.' as mp_position
		from category
		where parent_id is null

		union all

		select
			c.id,
			category_cte.name || ' > ' || c.name,
			category_cte.slug || c.slug || '/',
			c.attribute_group_id,
			category_cte.mp_position || printf('%04d', c.mp_position) || '.'
		from category c
		join category_cte on category_cte.id = c.parent_id
	)
`

func (st *Category) getQuerySelectBase(extra ...string) *sqlbuilder.SelectBuilder {
	columns := []string{
		"c.id",
		"c.name",
		"c.icon",
		"c.is_published",
		"c.parent_id",
		"pc.name as parent_text",
		"c.attribute_group_id",
		"ag.name as attribute_group_text",
		"c_cte.name as hierarchy",
		"c_cte.slug as category_url",
	}

	sb := sqlbuilder.NewSelectBuilder()
	sb.SQL(CategoryCTE)
	sb.Select(append(columns, extra...)...).From("category as c")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category_cte as c_cte", "c_cte.id = c.id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "category as pc", "pc.id = c.parent_id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "attribute_group as ag", "ag.id = c.attribute_group_id")
	return sb
}

func (st *Category) getQueryFindByParentID(parentID int64, q storage.FindManyParams) (string, []any) {
	sb := st.getQuerySelectBase(
		"coalesce(p.count, 0) as related_products",
		"count(*) over () as count",
	)

	sb.JoinWithOption(sqlbuilder.LeftJoin, `(
		select p.category_id, count(*) as count
		from product as p
		join supplier as s on s.id = p.supplier_id
		where p.quantity > 0 and p.is_published = 1 and s.is_published = 1
		group by p.category_id
	) as p`, "p.category_id = c.id")

	if parentID == 0 {
		sb.Where(sb.IsNull("c.parent_id"))
	} else {
		sb.Where(sb.EQ("c.parent_id", parentID))
	}

	allowedFilters := map[string]string{
		"name":             "c.name",
		"icon":             "c.icon",
		"attribute_group":  "coalesce(c.attribute_group_id, 0)",
		"is_published":     "c.is_published",
		"related_products": "coalesce(p.count, 0)",
	}

	storage.ApplyFilters(sb, q.Filters, q.LogicAnd, allowedFilters)
	storage.ApplyLimitOffset(sb, q.Limit, q.Offset)
	sb.OrderBy("c.mp_position", "c.id DESC")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Category) FindManyByParentID(ctx context.Context, parentID int64, q storage.FindManyParams) ([]model.Category, int64, error) {
	const op = "sqlitedb.Category.FindManyByParentID"

	type LocalResult struct {
		model.Category
		Count int64 `db:"count"`
	}

	query, args := st.getQueryFindByParentID(parentID, q)
	rows, err := runSelect[LocalResult](ctx, st.db, query, args)
	if err != nil {
		return nil, 0, utils.WrapIfErr(op, err)
	}

	dto, count := make([]model.Category, len(rows)), int64(0)
	for i, row := range rows {
		dto[i], count = row.Category, row.Count
	}

	return dto, count, nil
}

func (st *Category) FillCategoryChildren(ctx context.Context, category *model.Category, recursive bool) error {
	const op = "sqlitedb.Category.FillCategoryChildren"

	children, _, err := st.FindManyByParentID(ctx, category.ID, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	} else {
		category.Children = children
	}

	// recursive end condition
	if children == nil || !recursive {
		return nil
	}

	for i := range category.Children {
		if err := st.FillCategoryChildren(ctx, &category.Children[i], recursive); err != nil {
			return utils.WrapIfErr(op, err)
		}
	}

	return nil
}

func (st *Category) GetTree(ctx context.Context) (model.Category, error) {
	const op = "sqlitedb.Category.GetTree"

	root := &model.Category{Name: "All categories"}
	if err := st.FillCategoryChildren(ctx, root, true); err != nil {
		return model.Category{}, utils.WrapIfErr(op, err)
	}

	return *root, nil
}

func (st *Category) getQueryGetDropdown(search string, max int, leafsOnly bool) (string, []any) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SQL(CategoryCTE)
	sb.Select("c_cte.id", "c_cte.name as text", "c_cte.attribute_group_id")
	sb.From("category_cte as c_cte")

	if search != "" {
		sb.Where(sb.Like("c_cte.name", "%"+search+"%"))
	}

	if leafsOnly {
		sb.JoinWithOption(sqlbuilder.LeftJoin, "category as c_parent", "c_parent.parent_id = c_cte.id")
		sb.Where("c_parent.id is null")
	}

	sb.OrderBy("c_cte.mp_position").Limit(max)

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Category) GetDropdown(ctx context.Context, search string, max int, leafsOnly bool) ([]model.CategoryDropdown, error) {
	const op = "sqlitedb.Category.GetDropdown"
	query, args := st.getQueryGetDropdown(search, max, leafsOnly)
	rows, err := runSelect[model.CategoryDropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *Category) getQueryGetByID(id int64) (string, []any) {
	sb := st.getQuerySelectBase()
	sb.Where(sb.EQ("c.id", id))
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Category) GetByID(ctx context.Context, id int64) (model.Category, bool, error) {
	const op = "sqlitedb.Category.GetByID"
	query, args := st.getQueryGetByID(id)
	row, ok, err := runGet[model.Category](ctx, st.db, query, args)
	return row, ok, utils.WrapIfErr(op, err)
}

func (st *Category) getQueryDeleteManyByID(ids []int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("category")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Category) DeleteManyByID(ctx context.Context, ids []int64) (int64, error) {
	const op = "sqlitedb.Category.DeleteManyByID"
	query, args := st.getQueryDeleteManyByID(ids)
	affected, err := runExecAffected(ctx, st.db, query, args)
	return affected, utils.WrapIfErr(op, err)
}

func (st *Category) getQueryInsert(dto model.Category) (string, []any) {
	ib := sqlbuilder.InsertInto("category")
	ib.Cols("slug", "name", "icon", "is_published", "attribute_group_id", "parent_id")
	ib.Values(dto.Slugify(), dto.Name, dto.Icon, dto.IsPublished, dto.AttributeGroupID, dto.ParentID)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Category) getQueryUpdateByID(dto model.Category) (string, []any) {
	ub := sqlbuilder.Update("category")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	slug := dto.Slugify()
	if slug != "" {
		ub.SetMore(ub.EQ("slug", slug))
	}
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Icon", "icon", dto.Icon)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "IsPublished", "is_published", dto.IsPublished)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "CategoryParentEmbed.ParentID", "parent_id", dto.ParentID)
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "CategoryAttributeGroupEmbed.AttributeGroupID", "attribute_group_id", dto.AttributeGroupID)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Category) UpsertMany(ctx context.Context, dtos []model.Category) (int64, error) {
	const op = "sqlitedb.Category.UpsertMany"

	tx, err := st.db.Beginx()
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}
	defer tx.Rollback()

	var total int64

	for _, dto := range dtos {
		if dto.ID == 0 {
			_, err := st.insert(ctx, tx, dto)
			if err != nil {
				return 0, utils.WrapIfErr(op, err)
			}
			total += 1
		} else {
			affected, err := st.update(ctx, tx, dto)
			if err != nil {
				return 0, utils.WrapIfErr(op, err)
			}
			total += affected
		}
	}

	return total, utils.WrapIfErr(op, tx.Commit())
}

func (st *Category) insert(ctx context.Context, tx *sqlx.Tx, dto model.Category) (int64, error) {
	const op = "sqlitedb.Category.insert"

	query, args := st.getQueryInsert(dto)
	id, err := runExecInsert(ctx, tx, query, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	if err := st.updateMaterializedPath(ctx, tx, id, dto.ParentID); err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	return id, nil
}

func (st *Category) update(ctx context.Context, tx *sqlx.Tx, dto model.Category) (int64, error) {
	const op = "sqlitedb.Category.update"

	query, args := st.getQueryUpdateByID(dto)
	affected, err := runExecAffected(ctx, tx, query, args)
	if err != nil {
		return 0, utils.WrapIfErr(op, err)
	}

	if _, ok := dto.Partial["CategoryParentEmbed.ParentID"]; ok {
		ok, err := st.validateCategoryTree(ctx, tx, dto.ID, make(map[int64]struct{}))
		if err != nil {
			return 0, utils.WrapIfErr(op, err)
		} else if !ok {
			return 0, utils.WrapIfErr(op, storage.ErrInvalidTreeCircular)
		}

		if err := st.updateMaterializedPath(ctx, tx, dto.ID, dto.ParentID); err != nil {
			return 0, utils.WrapIfErr(op, err)
		}
	}

	return affected, nil
}

func (st *Category) validateCategoryTree(ctx context.Context, tx *sqlx.Tx, id int64, traversed map[int64]struct{}) (bool, error) {
	const op = "sqlitedb.Category.validateCategoryTree"

	// root node
	if id == 0 {
		return true, nil
	}

	// node has been already visited (circular)
	if _, ok := traversed[id]; ok {
		return false, nil
	}

	// visit the node
	traversed[id] = struct{}{}

	// select the parent node
	parentID, _, err := runGet[int64](ctx, tx, "select coalesce(parent_id, 0) from category where id = ?;", []any{id})
	if err != nil {
		return false, utils.WrapIfErr(op, err)
	}

	return st.validateCategoryTree(ctx, tx, parentID, traversed)
}

func (st *Category) updateMaterializedPath(ctx context.Context, tx *sqlx.Tx, id int64, parentID *int64) error {
	const op = "sqlitedb.Category.updateMaterializedPath"

	// step 1. update mp_path = parent path + new id
	const updatePathSQL = "update category set mp_path = coalesce( (select mp_path from category where id = ?) || ? || '.', ? || '.' ) where id = ?;"
	if _, err := runExecAffected(ctx, tx, updatePathSQL, []any{parentID, id, id, id}); err != nil {
		return utils.WrapIfErr(op, err)
	}

	// step 2. update mp_level = count of '.' in the path
	const updateLevelSQL = "update category set mp_level = length(mp_path) - length(replace(mp_path, '.', '')) - 1 where id = ?;"
	if _, err := runExecAffected(ctx, tx, updateLevelSQL, []any{id}); err != nil {
		return utils.WrapIfErr(op, err)
	}

	// step 3. set materialized path for child nodes
	type LocalChild struct {
		ID       int64  `db:"id"`
		ParentID *int64 `db:"parent_id"`
	}

	const selectChildsSQL = "select id, parent_id from category where parent_id = ?;"
	childs, err := runSelect[LocalChild](ctx, tx, selectChildsSQL, []any{id})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	for _, child := range childs {
		st.updateMaterializedPath(ctx, tx, child.ID, child.ParentID)
	}

	return nil
}

func (st *Category) getQueryUpdatePosition(id int64, position int64) (string, []any) {
	ub := sqlbuilder.Update("category")
	ub.Where(ub.EQ("id", id))
	ub.SetMore("updated_at = unixepoch()")
	ub.SetMore(ub.EQ("mp_position", position))
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Category) UpdatePositions(ctx context.Context, orderedIDs []int64) (int64, error) {
	const op = "sqlitedb.Category.UpdatePositions"

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
