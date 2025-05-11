package attribute

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

func (repo *SQLiteRepository) FindAllGroups(ctx context.Context) ([]AttributeGroup, error) {
	const op = "attribute.SQLiteRepository.FindAllGroups"

	var records []AttributeGroup

	sb := sqlbuilder.Select("id", "name")
	sb.From("attribute_group")
	sb.OrderBy("name")

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record AttributeGroup
		if err := rows.Scan(&record.ID, &record.Name); err != nil {
			return nil, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetGroupDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "attribute.SQLiteRepository.GetGroupDropdown"

	var records []w2.DropdownValue

	sb := sqlbuilder.Select("id", "name as text")
	sb.From("attribute_group")
	if r.Search != "" {
		sb.Where(sb.Like("name", "%"+r.Search+"%"))
	}
	sb.OrderBy("name")
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

func (repo *SQLiteRepository) UpsertGroup(ctx context.Context, dto AttributeGroup) error {
	const op = "attribute.SQLiteRepository.UpsertGroup"

	var query string
	var args []any

	if dto.ID == 0 {
		ib := sqlbuilder.InsertInto("attribute_group")
		ib.Cols("name")
		ib.Values(dto.Name)
		query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
	} else {
		ub := sqlbuilder.Update("attribute_group")
		ub.Where(ub.EQ("id", dto.ID))
		ub.SetMore("updated_at = unixepoch()")
		ub.SetMore(ub.EQ("name", dto.Name))
		query, args = ub.BuildWithFlavor(sqlbuilder.SQLite)
	}

	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) DeleteGroupByID(ctx context.Context, id int) error {
	const op = "attribute.SQLiteRepository.DeleteGroupByID"
	dlb := sqlbuilder.DeleteFrom("attribute_group")
	dlb.Where(dlb.EQ("id", id))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) FindManySetsByGroupID(ctx context.Context, groupID int, r w2.GridDataRequest) ([]AttributeSet, int, error) {
	const op = "attribute.SQLiteRepository.FindManySets"

	var total int
	var records []AttributeSet

	sb := sqlbuilder.Select("count(*)")
	sb.From("attribute_set")
	sb.Where(sb.EQ("attribute_group_id", groupID))

	w2sqlbuilder.Where(sb, r, map[string]string{
		"name": "name",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
		"id",
		"attribute_group_id",
		"name",
	)

	sb.OrderBy("position", "id DESC")

	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record AttributeSet
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetGroupIDBySetID(ctx context.Context, setID int) (int, error) {
	const op = "attribute.SQLiteRepository.GetGroupIDBySetID"
	var groupID int
	row := repo.store.DB().QueryRowContext(ctx, "select attribute_group_id from attribute_set where id = ?;", setID)
	return groupID, wrap.IfErr(op, row.Scan(&groupID))
}

func (repo *SQLiteRepository) UpsertManySets(ctx context.Context, groupID int, changes []AttributeSet) error {
	const op = "attribute.SQLiteRepository.UpsertManySets"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		var query string
		var args []any

		if dto.ID == 0 {
			ib := sqlbuilder.InsertInto("attribute_set")
			ib.Cols("attribute_group_id", "name")
			ib.Values(groupID, dto.Name)
			query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
		} else {
			ub := sqlbuilder.Update("attribute_set")
			ub.Where(ub.EQ("id", dto.ID))
			ub.SetMore("updated_at = unixepoch()")
			w2sqlbuilder.SetEditable(ub, dto.Name, "name")
			query, args = ub.BuildWithFlavor(sqlbuilder.SQLite)
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManySetsByID(ctx context.Context, ids []int) error {
	const op = "attribute.SQLiteRepository.DeleteManySetsByID"
	dlb := sqlbuilder.DeleteFrom("attribute_set")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) UpdateSetPositions(ctx context.Context, orderedIDs []int) error {
	const op = "attribute.SQLiteRepository.UpdateSetPositions"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for i, id := range orderedIDs {
		ub := sqlbuilder.Update("attribute_set")
		ub.Where(ub.EQ("id", id))
		ub.SetMore("updated_at = unixepoch()")
		ub.SetMore(ub.EQ("position", i))
		query, args := ub.BuildWithFlavor(sqlbuilder.SQLite)

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) FindManyValuesBySetID(ctx context.Context, setID int, r w2.GridDataRequest) ([]AttributeValue, int, error) {
	const op = "attribute.SQLiteRepository.FindManyValuesBySetID"

	var total int
	var records []AttributeValue

	sb := sqlbuilder.Select("count(*)")
	sb.From("attribute_value as av")

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

	w2sqlbuilder.Where(sb, r, map[string]string{
		"name": "av.name",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
		"av.id",
		"av.attribute_set_id",
		"av.name",
		"coalesce(p.related_products, 0) as related_products",
		"coalesce(p.published_products, 0) as published_products",
	)

	sb.OrderBy("av.position", "av.id DESC")

	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record AttributeValue
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetValueDropdownBySetID(ctx context.Context, setID int, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "attribute.SQLiteRepository.GetValueDropdownBySetID"

	var records []w2.DropdownValue

	sb := sqlbuilder.Select("id", "name as text")
	sb.From("attribute_value")
	sb.Where(sb.EQ("attribute_set_id", setID))
	if r.Search != "" {
		sb.Where(sb.Like("name", "%"+r.Search+"%"))
	}
	sb.OrderBy("position", "id DESC")
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

func (repo *SQLiteRepository) GetSetIDByValueID(ctx context.Context, valueID int) (int, error) {
	const op = "attribute.SQLiteRepository.GetSetIDByValueID"
	var setID int
	row := repo.store.DB().QueryRowContext(ctx, "select attribute_set_id from attribute_value where id = ?;", valueID)
	return setID, wrap.IfErr(op, row.Scan(&setID))
}

func (repo *SQLiteRepository) UpsertManyValues(ctx context.Context, setID int, changes []AttributeValue) error {
	const op = "attribute.SQLiteRepository.UpsertManyValues"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		var query string
		var args []any

		if dto.ID == 0 {
			ib := sqlbuilder.InsertInto("attribute_value")
			ib.Cols("attribute_set_id", "name")
			ib.Values(setID, dto.Name)
			query, args = ib.BuildWithFlavor(sqlbuilder.SQLite)
		} else {
			ub := sqlbuilder.Update("attribute_value")
			ub.Where(ub.EQ("id", dto.ID))
			ub.SetMore("updated_at = unixepoch()")
			w2sqlbuilder.SetEditable(ub, dto.Name, "name")
			query, args = ub.BuildWithFlavor(sqlbuilder.SQLite)
		}

		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManyValuesByID(ctx context.Context, ids []int) error {
	const op = "attribute.SQLiteRepository.DeleteManyValuesByID"
	dlb := sqlbuilder.DeleteFrom("attribute_value")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) UpdateValuePositions(ctx context.Context, orderedIDs []int) error {
	const op = "attribute.SQLiteRepository.UpdateValuePositions"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for i, id := range orderedIDs {
		ub := sqlbuilder.Update("attribute_value")
		ub.Where(ub.EQ("id", id))
		ub.SetMore("updated_at = unixepoch()")
		ub.SetMore(ub.EQ("position", i))

		query, args := ub.BuildWithFlavor(sqlbuilder.SQLite)
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}
