package sqlitedb

import (
	"context"
	"fmt"
	"maps"
	"strconv"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type Filter struct {
	db *sqlx.DB
}

func NewFilter(db *sqlx.DB) *Filter {
	return &Filter{db: db}
}

func (st *Filter) getQueryFindFiltersByCategoryID(categoryID int64) (string, []any) {
	sb0 := sqlbuilder.Select("'S' as id", "'Supplier' as name", "-1 as position")
	sb1 := sqlbuilder.Select("'B' as id", "'Brand' as name", "-2 as position")

	sb2 := sqlbuilder.Select("ats.id", "ats.name", "ats.position")
	sb2.From("attribute_set as ats")
	sb2.JoinWithOption(sqlbuilder.InnerJoin, "attribute_group as atg", "atg.id = ats.attribute_group_id")
	sb2.JoinWithOption(sqlbuilder.InnerJoin, "category as c", "c.attribute_group_id = atg.id")
	sb2.Where(sb2.EQ("c.id", categoryID), "ats.in_filter = 1")

	b := sqlbuilder.Build("select id, name from ($?) as t1", sqlbuilder.Union(sb0, sb1, sb2).OrderBy("position"))
	return b.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Filter) FindFiltersByCategoryID(ctx context.Context, categoryID int64, applied model.FilterCombination) ([]model.Filter, error) {
	const op = "sqlitedb.Filter.FindFiltersByCategoryID"
	query, args := st.getQueryFindFiltersByCategoryID(categoryID)
	rows, err := runSelect[model.Filter](ctx, st.db, query, args)
	for i := range rows {
		if err := st.fillFilterFacets(ctx, &rows[i], categoryID, applied); err != nil {
			return rows, utils.WrapIfErr(op, err)
		}
	}
	return rows, utils.WrapIfErr(op, err)
}

func (st *Filter) fillFilterFacets(ctx context.Context, filter *model.Filter, categoryID int64, applied model.FilterCombination) error {
	const op = "sqlitedb.Filter.fillFilterFacets"

	var facets []model.FilterFacet
	var err error

	if filter.ID == "S" {
		query, args := st.getQueryFindSupplierFacets(categoryID, applied)
		facets, err = runSelect[model.FilterFacet](ctx, st.db, query, args)
	} else if filter.ID == "B" {
		query, args := st.getQueryFindBrandFacets(categoryID, applied)
		facets, err = runSelect[model.FilterFacet](ctx, st.db, query, args)
	} else {
		query, args := st.getQueryFindAttributeFacets(categoryID, filter.ID, applied)
		facets, err = runSelect[model.FilterFacet](ctx, st.db, query, args)
	}

	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	for i := range facets {
		// if there is something selected in this filter group
		if len(applied[filter.ID]) > 0 {
			// if current value is selected
			if _, ok := applied[filter.ID][facets[i].ID]; ok {
				facets[i].IsSelected = true
			}
		}
	}

	filter.Facets = facets
	filter.Selected = len(applied[filter.ID])

	return nil
}

func (st *Filter) getQueryFindSupplierFacets(selectedCategoryID int64, applied model.FilterCombination) (string, []any) {
	facet := sqlbuilder.Select("p.id").From("product as p")
	for filterKey, facets := range applied {
		if filterKey == "S" {
			continue
		} else if filterKey == "B" {
			facet.Where(facet.In("p.brand_id", sqlbuilder.List(maps.Keys(facets))))
		} else if attributeSetID, err := strconv.Atoi(filterKey); err == nil {
			joinTable := fmt.Sprintf("product_attribute as pa_%d", attributeSetID)
			joinExpr1 := fmt.Sprintf("pa_%d.product_id = p.id", attributeSetID)
			joinExpr2 := facet.In(fmt.Sprintf("pa_%d.attribute_value_id", attributeSetID), sqlbuilder.List(maps.Keys(facets)))
			facet.JoinWithOption(sqlbuilder.InnerJoin, joinTable, facet.And(joinExpr1, joinExpr2))
		}
	}

	sb := sqlbuilder.Select("s.id", "s.name", "count(facet.id) as products")
	sb.From("product as p")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "supplier as s", "s.id = p.supplier_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "brand as b", "b.id = p.brand_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c", "c.id = p.category_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c_parent", "c.mp_path like c_parent.mp_path || '%'")
	sb.JoinWithOption(sqlbuilder.LeftJoin, sb.Var(sqlbuilder.Buildf("(%v) as facet", facet)), "facet.id = p.id")
	sb.Where(sb.EQ("c_parent.id", selectedCategoryID), "p.quantity > 0 and p.is_published = 1 and s.is_published = 1")
	sb.GroupBy("s.id", "s.name")
	sb.OrderBy("s.position")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Filter) getQueryFindBrandFacets(selectedCategoryID int64, applied model.FilterCombination) (string, []any) {
	facet := sqlbuilder.Select("p.id").From("product as p")
	for filterKey, facets := range applied {
		if filterKey == "B" {
			continue
		} else if filterKey == "S" {
			facet.Where(facet.In("p.supplier_id", sqlbuilder.List(maps.Keys(facets))))
		} else if attributeSetID, err := strconv.Atoi(filterKey); err == nil {
			joinTable := fmt.Sprintf("product_attribute as pa_%d", attributeSetID)
			joinExpr1 := fmt.Sprintf("pa_%d.product_id = p.id", attributeSetID)
			joinExpr2 := facet.In(fmt.Sprintf("pa_%d.attribute_value_id", attributeSetID), sqlbuilder.List(maps.Keys(facets)))
			facet.JoinWithOption(sqlbuilder.InnerJoin, joinTable, facet.And(joinExpr1, joinExpr2))
		}
	}

	sb := sqlbuilder.Select("b.id", "b.name", "count(facet.id) as products")
	sb.From("product as p")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "supplier as s", "s.id = p.supplier_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "brand as b", "b.id = p.brand_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c", "c.id = p.category_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c_parent", "c.mp_path like c_parent.mp_path || '%'")
	sb.JoinWithOption(sqlbuilder.LeftJoin, sb.Var(sqlbuilder.Buildf("(%v) as facet", facet)), "facet.id = p.id")
	sb.Where(sb.EQ("c_parent.id", selectedCategoryID), "p.quantity > 0 and p.is_published = 1 and s.is_published = 1")
	sb.GroupBy("b.id", "b.name")
	sb.OrderBy("b.name")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Filter) getQueryFindAttributeFacets(selectedCategoryID int64, selectedFilterKey string, applied model.FilterCombination) (string, []any) {
	facet := sqlbuilder.Select("p.id").From("product as p")
	for filterKey, facets := range applied {
		if filterKey == selectedFilterKey {
			continue
		} else if filterKey == "S" {
			facet.Where(facet.In("p.supplier_id", sqlbuilder.List(maps.Keys(facets))))
		} else if filterKey == "B" {
			facet.Where(facet.In("p.brand_id", sqlbuilder.List(maps.Keys(facets))))
		} else if attributeSetID, err := strconv.Atoi(filterKey); err == nil {
			joinTable := fmt.Sprintf("product_attribute as pa_%d", attributeSetID)
			joinExpr1 := fmt.Sprintf("pa_%d.product_id = p.id", attributeSetID)
			joinExpr2 := facet.In(fmt.Sprintf("pa_%d.attribute_value_id", attributeSetID), sqlbuilder.List(maps.Keys(facets)))
			facet.JoinWithOption(sqlbuilder.InnerJoin, joinTable, facet.And(joinExpr1, joinExpr2))
		}
	}

	sb := sqlbuilder.Select("atv.id", "atv.name", "count(facet.id) as products")
	sb.From("product as p")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "supplier as s", "s.id = p.supplier_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "product_attribute as pa", "pa.product_id = p.id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "attribute_value as atv", "atv.id = pa.attribute_value_id")
	sb.JoinWithOption(sqlbuilder.LeftJoin, sb.Var(sqlbuilder.Buildf("(%v) as facet", facet)), "facet.id = p.id")
	sb.Where(
		sb.EQ("p.category_id", selectedCategoryID),
		sb.EQ("pa.attribute_set_id", selectedFilterKey),
		"p.quantity > 0 and p.is_published = 1 and s.is_published = 1",
	)
	sb.GroupBy("atv.id", "atv.name")
	sb.OrderBy("atv.position")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}
