package product

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

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

func (repo *SQLiteRepository) selectBase() *sqlbuilder.SelectBuilder {
	sb := sqlbuilder.NewSelectBuilder()
	sb.SQL(`
		with recursive category_cte as (
			select
				id,
				name,
				slug || '/' as slug
			from category
			where parent_id is null

			union all

			select
				c.id,
				category_cte.name || ' > ' || c.name,
				category_cte.slug || c.slug || '/'
			from category c
			join category_cte on category_cte.id = c.parent_id
		)
	`)
	sb.From("product as p")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c", "c.id = p.category_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c_root", "c_root.mp_level = 0 and c.mp_path like c_root.mp_path || '%'")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category_cte as c_cte", "c_cte.id = p.category_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "supplier as s", "s.id = p.supplier_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "brand as b", "b.id = p.brand_id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "product_status as ps", "ps.id = p.status_id")
	return sb
}

func (repo *SQLiteRepository) selectColumns(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
	return sb.Select(
		"p.id",
		"p.code",
		"p.name",
		"p.description",
		"p.quantity",
		"p.price",
		"p.is_published",
		"datetime(p.created_at, 'unixepoch', 'localtime') as created_at",
		"datetime(p.updated_at, 'unixepoch', 'localtime') as updated_at",
		"p.category_id",
		"c_cte.name as category_text",
		"p.supplier_id",
		"s.name as supplier_text",
		"p.brand_id",
		"b.name as brand_text",
		"p.status_id",
		"ps.name as status_text",
		"ps.color as status_color",
		"concat(s.slug, '/', p.slug, '/') as product_url",
	)
}

func (repo *SQLiteRepository) selectAvailable(sb *sqlbuilder.SelectBuilder) *sqlbuilder.SelectBuilder {
	sb.Where("p.quantity > 0 and p.is_published = 1 and c.is_published = 1 and s.is_published = 1")
	return sb
}

func (repo *SQLiteRepository) selectByCategory(sb *sqlbuilder.SelectBuilder, categoryID int) *sqlbuilder.SelectBuilder {
	if categoryID != 0 {
		sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c_selected", "c.mp_path like c_selected.mp_path || '%'")
		sb.Where(sb.EQ("c_selected.id", categoryID))
	}
	return sb
}

func (repo *SQLiteRepository) FindManyByCategoryID(ctx context.Context, r w2.GridDataRequest, categoryID int) ([]Product, int, error) {
	const op = "product.SQLiteRepository.FindMany"

	var total int
	var records []Product

	sb := repo.selectBase()
	repo.selectByCategory(sb, categoryID)
	sb.Select("count(*)")

	for _, v := range r.Search {
		if v.Field == "attribute" {
			sb.JoinWithOption(sqlbuilder.InnerJoin, "product_attribute as pa on pa.product_id = p.id")
			sb.Where(sb.EQ("pa.attribute_value_id", v.Value))
			break
		}
	}

	w2sqlbuilder.Where(sb, r, map[string]string{
		"code":         "p.code",
		"name":         "p.name",
		"category":     "p.category_id",
		"supplier":     "p.supplier_id",
		"brand":        "p.brand_id",
		"status":       "coalesce(p.status_id, 0)",
		"is_published": "p.is_published",
		"quantity":     "p.quantity",
		"price":        "p.price",
	})

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	repo.selectColumns(sb)
	w2sqlbuilder.OrderBy(sb, r, map[string]string{
		"id":           "p.id",
		"code":         "p.code",
		"name":         "p.name",
		"supplier":     "s.name",
		"brand":        "b.name",
		"status":       "ps.name",
		"quantity":     "p.quantity",
		"price":        "p.price",
		"is_published": "p.is_published",
		"created_at":   "p.created_at",
		"updated_at":   "p.updated_at",
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
		var record Product
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}

		if err := repo.fillDtoDetails(ctx, &record); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}

		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) FindAvailableByCategoryID(ctx context.Context, r ListDataRequest, categoryID int) ([]Product, error) {
	const op = "product.SQLiteRepository.FindAvailableByCategoryID"

	var records []Product

	sb := repo.selectBase()
	repo.selectColumns(sb)
	repo.selectAvailable(sb)
	repo.selectByCategory(sb, categoryID)

	if r.Search != "" {
		switch r.SearchBy {
		case "Code":
			sb.Where(sb.Like("p.code", "%"+r.Search+"%"))
		case "Name":
			sb.Where(sb.Like("p.name", "%"+r.Search+"%"))
		}
	}

	if r.Filters != "" {
		var filters map[string][]string
		if err := json.Unmarshal([]byte(r.Filters), &filters); err != nil {
			return nil, err
		}

		mapping := map[string]string{
			"status": "p.status_id",
		}

		for key, values := range filters {
			sb.Where(sb.In(mapping[key], sqlbuilder.List(values)))
		}
	}

	sb.Limit(r.Limit)
	sb.Offset(r.Offset)

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record Product
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, wrap.IfErr(op, err)
		}

		if err := repo.fillDtoDetails(ctx, &record); err != nil {
			return nil, wrap.IfErr(op, err)
		}

		records = append(records, record)
	}

	return records, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) GetByID(ctx context.Context, id int) (Product, error) {
	const op = "product.SQLiteRepository.GetByID"

	var product Product

	sb := repo.selectBase()
	repo.selectColumns(sb)
	sb.Where(sb.EQ("p.id", id))

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	return product, wrap.IfErr(op, product.ScanRow(row.Scan))
}

func (repo *SQLiteRepository) fillDtoDetails(ctx context.Context, dto *Product) error {
	const op = "product.SQLiteRepository.fillDtoDetails"

	var err error

	if dto.Attributes, _, err = repo.FindManyAttributesByProductID(ctx, dto.ID, w2.GridDataRequest{}); err != nil {
		return wrap.IfErr(op, err)
	}

	return nil
}

func (repo *SQLiteRepository) insertQuery(dto Product) (string, []any) {
	ib := sqlbuilder.InsertInto("product")
	ib.Cols(
		"code",
		"slug",
		"name",
		"description",
		"quantity",
		"price",
		"brand_id",
		"category_id",
		"supplier_id",
		"status_id",
		"is_published",
	)
	ib.Values(
		dto.Code,
		dto.Slugify(),
		dto.Name,
		dto.Description,
		dto.Quantity,
		dto.Price,
		dto.Brand.ID,
		dto.Category.ID,
		dto.Supplier.ID,
		dto.Status.ID,
		dto.IsPublished.V,
	)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (repo *SQLiteRepository) updateQuery(dto Product) (string, []any) {
	ub := sqlbuilder.Update("product")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	slug := dto.Slugify()
	if slug != "" {
		ub.SetMore(ub.EQ("slug", slug))
	}
	w2sqlbuilder.SetEditable(ub, dto.Code, "code")
	w2sqlbuilder.SetEditable(ub, dto.Name, "name")
	w2sqlbuilder.SetEditable(ub, dto.Description, "description")
	w2sqlbuilder.SetEditable(ub, dto.Quantity, "quantity")
	w2sqlbuilder.SetEditable(ub, dto.Price, "price")
	w2sqlbuilder.SetEditable(ub, dto.Brand.ID, "brand_id")
	w2sqlbuilder.SetEditable(ub, dto.Category.ID, "category_id")
	w2sqlbuilder.SetEditable(ub, dto.Supplier.ID, "supplier_id")
	w2sqlbuilder.SetEditable(ub, dto.Status.ID, "status_id")
	w2sqlbuilder.SetEditable(ub, dto.IsPublished, "is_published")
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (repo *SQLiteRepository) UpsertOne(ctx context.Context, dto Product) (int, error) {
	const op = "product.SQLiteRepository.UpsertOne"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return 0, wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	if dto.ID == 0 {
		query, args := repo.insertQuery(dto)
		res, err := tx.ExecContext(ctx, query, args...)
		if err != nil {
			return 0, wrap.IfErr(op, err)
		}
		id, err := res.LastInsertId()
		if err != nil {
			return 0, wrap.IfErr(op, err)
		}
		dto.ID = int(id)
	} else {
		query, args := repo.updateQuery(dto)
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return 0, wrap.IfErr(op, err)
		}
	}

	const deleteInvalidAttributesSQL = `
		delete from product_attribute
		where product_id = ? and attribute_set_id not in (
			select ats.id
			from product as p
			join category as c on c.id = p.category_id
			join attribute_group as atg on atg.id = c.attribute_group_id
			join attribute_set as ats on ats.attribute_group_id = atg.id
			where p.id = ?
		);
	`

	if _, err := tx.ExecContext(ctx, deleteInvalidAttributesSQL, dto.ID, dto.ID); err != nil {
		return 0, wrap.IfErr(op, fmt.Errorf("deleteInvalidAttributesSQL: %w", err))
	}

	return dto.ID, wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) UpdateMany(ctx context.Context, changes []Product) error {
	const op = "product.SQLiteRepository.UpdateMany"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		query, args := repo.updateQuery(dto)
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}

func (repo *SQLiteRepository) DeleteManyByID(ctx context.Context, ids []int) error {
	const op = "product.SQLiteRepository.DeleteManyByID"
	dlb := sqlbuilder.DeleteFrom("product")
	dlb.Where(dlb.In("id", sqlbuilder.List(ids)))
	query, args := dlb.BuildWithFlavor(sqlbuilder.SQLite)
	_, err := repo.store.DB().ExecContext(ctx, query, args...)
	return wrap.IfErr(op, err)
}

func (repo *SQLiteRepository) FindManyAttributesByProductID(ctx context.Context, productID int, r w2.GridDataRequest) ([]ProductAttribute, int, error) {
	const op = "product.SQLiteRepository.FindManyAttributesByProductID"

	var total int
	var records []ProductAttribute

	sb := sqlbuilder.Select("count(*)")
	sb.From("product as p")
	sb.Where(sb.EQ("p.id", productID))

	query, args := sb.BuildWithFlavor(sqlbuilder.SQLite)
	row := repo.store.DB().QueryRowContext(ctx, query, args...)
	if err := row.Scan(&total); err != nil && err != sql.ErrNoRows {
		return nil, 0, wrap.IfErr(op, err)
	}

	sb.Select(
		"ats.id as id",
		"ats.name",
		"atv.id as value_id",
		"atv.name as value_text",
	)

	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c", "c.id = p.category_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "attribute_group as atg", "atg.id = c.attribute_group_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "attribute_set as ats", "ats.attribute_group_id = atg.id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "product_attribute as pa", "pa.product_id = p.id and pa.attribute_set_id = ats.id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, "attribute_value as atv", "atv.id = pa.attribute_value_id")
	sb.OrderBy("ats.position", "ats.id DESC")

	w2sqlbuilder.Limit(sb, r)
	w2sqlbuilder.Offset(sb, r)

	query, args = sb.BuildWithFlavor(sqlbuilder.SQLite)
	rows, err := repo.store.DB().QueryContext(ctx, query, args...)
	if err != nil {
		return nil, 0, wrap.IfErr(op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var record ProductAttribute
		if err := record.ScanRow(rows.Scan); err != nil {
			return nil, 0, wrap.IfErr(op, err)
		}
		records = append(records, record)
	}

	return records, total, wrap.IfErr(op, rows.Err())
}

func (repo *SQLiteRepository) UpsertManyAttributes(ctx context.Context, productID int, changes []ProductAttribute) error {
	const op = "product.SQLiteRepository.UpsertManyAttributes"

	tx, err := repo.store.DB().Begin()
	if err != nil {
		return wrap.IfErr(op, err)
	}
	defer tx.Rollback()

	for _, dto := range changes {
		ib := sqlbuilder.ReplaceInto("product_attribute")
		ib.Cols("product_id", "attribute_set_id", "attribute_value_id")
		ib.Values(productID, dto.AttributeSetID, dto.AttributeValue.ID)

		query, args := ib.BuildWithFlavor(sqlbuilder.SQLite)
		if _, err := tx.ExecContext(ctx, query, args...); err != nil {
			return wrap.IfErr(op, err)
		}
	}

	return wrap.IfErr(op, tx.Commit())
}
