package category

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
	cache Category
}

func NewSQLiteRepository(store db.Store) *SQLiteRepository {
	return &SQLiteRepository{store: store}
}

func (repo *SQLiteRepository) Store() db.Store {
	return repo.store
}

const categoryCTE = `
	with recursive category_cte as (
		select
			id,
			name,
			slug || '/' as slug,
			printf('%04d', mp_position) || '.' as mp_position
		from category
		where parent_id is null

		union all

		select
			c.id,
			category_cte.name || ' > ' || c.name,
			category_cte.slug || c.slug || '/',
			category_cte.mp_position || printf('%04d', c.mp_position) || '.'
		from category c
		join category_cte on category_cte.id = c.parent_id
	)
`

func (repo *SQLiteRepository) FindManyByParentID(ctx context.Context, parentID int, r w2.GridDataRequest) ([]Category, int, error) {
	const op = "category.SQLiteRepository.FindManyByParentID"

	var total int
	var records []Category

	sb := sqlbuilder.NewSelectBuilder()
	sb.SQL(categoryCTE)
	sb.Select("count(*)")
	sb.From("category as c")

	sb.JoinWithOption(sqlbuilder.InnerJoin, "category_cte as c_cte", "c_cte.id = c.id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "category as pc", "pc.id = c.parent_id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "attribute_group as ag", "ag.id = c.attribute_group_id")
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

	w2sqlbuilder.Where(sb, r, map[string]string{
		"name":             "c.name",
		"icon":             "c.icon",
		"attribute_group":  "coalesce(c.attribute_group_id, 0)",
		"is_published":     "c.is_published",
		"related_products": "coalesce(p.count, 0)",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
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
		"coalesce(p.count, 0) as related_products",
	)

	sb.OrderBy("c.mp_position", "c.id DESC")

	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Category
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetTree(ctx context.Context, useCache bool) (Category, error) {
	const op = "category.SQLiteRepository.GetTree"

	if useCache && len(repo.cache.Children) > 0 {
		return repo.cache, nil
	}

	root := Category{Name: w2.NewEditableWithValue("All categories")}
	if err := repo.fillCategoryChildren(ctx, &root, true); err != nil {
		return root, wrap.IfErr(op, err)
	}

	repo.cache = root
	return root, nil
}

func (repo *SQLiteRepository) fillCategoryChildren(ctx context.Context, category *Category, recursive bool) error {
	const op = "category.SQLiteRepository.fillCategoryChildren"

	children, _, err := repo.FindManyByParentID(ctx, category.ID, w2.GridDataRequest{})
	if err != nil {
		return wrap.IfErr(op, err)
	}

	category.Children = children

	// recursive end condition
	if children == nil || !recursive {
		return nil
	}

	for i := range category.Children {
		if err := repo.fillCategoryChildren(ctx, &category.Children[i], recursive); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return nil
}

func (repo *SQLiteRepository) GetDropdown(ctx context.Context, r w2.DropdownRequest, leafsOnly bool) ([]w2.DropdownValue, error) {
	const op = "category.SQLiteRepository.GetDropdown"

	var records []w2.DropdownValue

	sb := sqlbuilder.NewSelectBuilder()
	sb.SQL(categoryCTE)
	sb.Select("c_cte.id", "c_cte.name as text")
	sb.From("category_cte as c_cte")
	sb.OrderBy("c_cte.mp_position")

	if leafsOnly {
		sb.JoinWithOption(sqlbuilder.LeftJoin, "category as c_parent", "c_parent.parent_id = c_cte.id")
		sb.Where("c_parent.id is null")
	}

	if r.Search != "" {
		sb.Where(sb.Like("c_cte.name", "%"+r.Search+"%"))
	}

	sb.Limit(r.Max)

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record w2.DropdownValue
		if err := rows.Scan(&record.ID, &record.Text); err != nil {
			return nil, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetParentIDByCategoryID(ctx context.Context, categoryID int) (int, error) {
	const op = "category.SQLiteRepository.GetParentIDByCategoryID"
	var parentID int
	row := repo.store.DB().QueryRowContext(ctx, "select coalesce(parent_id, 0) from category where id = ?;", categoryID)
	return parentID, wrap.IfErr(op, row.Scan(&parentID))
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "category.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("category")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	repo.cache = Category{} // invalidate cache
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) UpsertMany(ctx context.Context, changes []Category) error {
	const op = "category.SQLiteRepository.UpsertMany"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		if dto.ID == 0 {
			if dto.Parent.ID.V == 0 {
				// allow insert on zero level with w2ui defaults
				// show <<CURRENT>> in the cell and ID is 0
				dto.Parent.ID.Valid = false
			}

			ib := sqlbuilder.InsertInto("category")
			ib.Cols("slug", "name", "icon", "is_published", "attribute_group_id", "parent_id")
			ib.Values(dto.Slugify(), dto.Name, dto.Icon, dto.IsPublished.V, dto.AttributeGroup.ID, dto.Parent.ID)

			query, args := ib.BuildWithFlavor(sqlbuilder.SQLite)
			res, err := tx.ExecContext(ctx, query, args...)
			if err != nil {
				return wrap.IfErr(op, err)
			}

			id, err := res.LastInsertId()
			if err != nil {
				return wrap.IfErr(op, err)
			}

			if err = repo.updateMaterializedPath(ctx, tx, int(id), dto.Parent.ID.V); err != nil {
				return wrap.IfErr(op, err)
			}
		} else {
			ub := sqlbuilder.Update("category")
			ub.Where(ub.EQ("id", dto.ID))
			ub.SetMore("updated_at = unixepoch()")

			slug := dto.Slugify()
			if slug != "" {
				ub.SetMore(ub.EQ("slug", slug))
			}

			w2sqlbuilder.SetEditable(ub, dto.Name, "name")
			w2sqlbuilder.SetEditable(ub, dto.Icon, "icon")
			w2sqlbuilder.SetEditable(ub, dto.IsPublished, "is_published")
			w2sqlbuilder.SetEditable(ub, dto.Parent.ID, "parent_id")
			w2sqlbuilder.SetEditable(ub, dto.AttributeGroup.ID, "attribute_group_id")

			query, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
			if _, err := tx.ExecContext(ctx, query, args...); err != nil {
				return wrap.IfErr(op, err)
			}

			if dto.Parent.ID.Provided {
				if err := repo.validateCategoryTree(ctx, tx, dto.ID); err != nil {
					return wrap.IfErr(op, err)
				}

				if err = repo.updateMaterializedPath(ctx, tx, dto.ID, dto.Parent.ID.V); err != nil {
					return wrap.IfErr(op, err)
				}
			}
		}
	}

	repo.cache = Category{} // invalidate cache
	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) UpdatePositions(ctx context.Context, orderedIDs []int) error {
	const op = "category.SQLiteRepository.UpdatePositions"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for i, id := range orderedIDs {
		ub := sqlbuilder.Update("category")
		ub.Where(ub.EQ("id", id))
		ub.SetMore("updated_at = unixepoch()")
		ub.SetMore(ub.EQ("mp_position", i))

		query, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	repo.cache = Category{} // invalidate cache
	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) updateMaterializedPath(ctx context.Context, tx *sql.Tx, id int, parentID int) error {
	const op = "category.SQLiteRepository.updateMaterializedPath"

	// step 1. update mp_path = parent path + new id
	const updatePathSQL = "update category set mp_path = coalesce( (select mp_path from category where id = ?) || ? || '.', ? || '.' ) where id = ?;"
	if _, err := tx.ExecContext(ctx, updatePathSQL, parentID, id, id, id); err != nil {
		return wrap.IfErr(op, err)
	}

	// step 2. update mp_level = count of '.' in the path
	const updateLevelSQL = "update category set mp_level = length(mp_path) - length(replace(mp_path, '.', '')) - 1 where id = ?;"
	if _, err := tx.ExecContext(ctx, updateLevelSQL, id); err != nil {
		return wrap.IfErr(op, err)
	}

	// step 3. set materialized path for child nodes
	const selectChildsSQL = "select id, parent_id from category where parent_id = ?;"
	rows, err := repo.store.DB().QueryContext(ctx, selectChildsSQL, id)
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var childCategoryID, childParentID int
		if err := rows.Scan(&childCategoryID, &childParentID); err != nil {
			return wrap.IfErr(op, err)
		}

		if err := repo.updateMaterializedPath(ctx, tx, childCategoryID, childParentID); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) validateCategoryTree(ctx context.Context, tx *sql.Tx, id int) error {
	const op = "category.SQLiteRepository.validateCategoryTree"
	traversed := map[int]struct{}{}

	for id != 0 {
		// Check if the node has already been visited
		if _, ok := traversed[id]; ok {
			return ErrCategoryInvalidCircular
		}

		// Mark the node as visited
		traversed[id] = struct{}{}

		// Fetch the parent ID
		var parentID int
		const fetchParentSQL = "select coalesce(parent_id, 0) from category where id = ?;"
		row := tx.QueryRowContext(ctx, fetchParentSQL, id)
		if err := row.Scan(&parentID); err != nil && err != sql.ErrNoRows {
			return wrap.IfErr(op, err)
		}

		// Move to the parent node
		id = parentID
	}

	return nil
}
