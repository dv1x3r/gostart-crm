package sqlitedb

import (
	"context"
	"fmt"
	"maps"
	"slices"
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

func (st *Filter) getQueryFindFacetsByCategoryID(categoryID int64) (string, []any) {
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

func (st *Filter) FindFacetsByCategoryID(ctx context.Context, categoryID int64, applied model.FilterCombination) ([]model.FilterFacet, error) {
	const op = "sqlitedb.Filter.FindFacetsByCategoryID"

	query, args := st.getQueryFindFacetsByCategoryID(categoryID)
	rows, err := runSelect[model.FilterFacet](ctx, st.db, query, args)
	if err != nil {
		return rows, utils.WrapIfErr(op, err)
	}

	for i := range rows {
		if err := st.fillFilterFacetsByCategoryID(ctx, &rows[i], categoryID, applied); err != nil {
			return rows, utils.WrapIfErr(op, err)
		}
	}

	return rows, utils.WrapIfErr(op, err)
}

func (st *Filter) fillFilterFacetsByCategoryID(ctx context.Context, facet *model.FilterFacet, categoryID int64, applied model.FilterCombination) error {
	const op = "sqlitedb.Filter.fillFilterFacetsByCategoryID"

	var values []model.FilterFacetValue
	var err error

	if facet.ID == "S" {
		query, args := st.getQueryFindSupplierFacetsByCategoryID(categoryID, applied)
		values, err = runSelect[model.FilterFacetValue](ctx, st.db, query, args)
	} else if facet.ID == "B" {
		query, args := st.getQueryFindBrandFacetsByCategoryID(categoryID, applied)
		values, err = runSelect[model.FilterFacetValue](ctx, st.db, query, args)
	} else {
		query, args := st.getQueryFindAttributeFacetsByCategoryID(categoryID, facet.ID, applied)
		values, err = runSelect[model.FilterFacetValue](ctx, st.db, query, args)
	}

	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	for i := range values {
		// if there is something selected in this filter group
		if len(applied[facet.ID]) > 0 {
			// if current value is selected
			if _, ok := applied[facet.ID][values[i].ID]; ok {
				values[i].IsSelected = true
			}
		}
	}

	facet.Values = values
	facet.Selected = len(applied[facet.ID])

	return nil
}

func (st *Filter) getQueryFindSupplierFacetsByCategoryID(categoryID int64, applied model.FilterCombination) (string, []any) {
	facet := sqlbuilder.Select("p.id").From("product as p")
	for filterKey, facets := range applied {
		facetKeys := slices.Collect(maps.Keys(facets))
		if filterKey == "S" {
			continue
		} else if filterKey == "B" {
			facet.Where(facet.In("p.brand_id", sqlbuilder.List(facetKeys)))
		} else if filterKey == "PF" && len(facetKeys) == 1 {
			facet.Where(facet.GTE("p.web_actual", facetKeys[0]))
		} else if filterKey == "PT" && len(facetKeys) == 1 {
			facet.Where(facet.LTE("p.web_actual", facetKeys[0]))
		} else if attributeSetID, err := strconv.Atoi(filterKey); err == nil {
			joinTable := fmt.Sprintf("product_attribute as pa_%d", attributeSetID)
			joinExpr1 := fmt.Sprintf("pa_%d.product_id = p.id", attributeSetID)
			joinExpr2 := facet.In(fmt.Sprintf("pa_%d.attribute_value_id", attributeSetID), sqlbuilder.List(facetKeys))
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
	sb.Where(sb.EQ("c_parent.id", categoryID), "p.quantity > 0 and p.is_published = 1 and s.is_published = 1")
	sb.GroupBy("s.id", "s.name")
	sb.OrderBy("s.position")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Filter) getQueryFindBrandFacetsByCategoryID(categoryID int64, applied model.FilterCombination) (string, []any) {
	facet := sqlbuilder.Select("p.id").From("product as p")
	for filterKey, facets := range applied {
		facetKeys := slices.Collect(maps.Keys(facets))
		if filterKey == "B" {
			continue
		} else if filterKey == "S" {
			facet.Where(facet.In("p.supplier_id", sqlbuilder.List(facetKeys)))
		} else if filterKey == "PF" && len(facetKeys) == 1 {
			facet.Where(facet.GTE("p.web_actual", facetKeys[0]))
		} else if filterKey == "PT" && len(facetKeys) == 1 {
			facet.Where(facet.LTE("p.web_actual", facetKeys[0]))
		} else if attributeSetID, err := strconv.Atoi(filterKey); err == nil {
			joinTable := fmt.Sprintf("product_attribute as pa_%d", attributeSetID)
			joinExpr1 := fmt.Sprintf("pa_%d.product_id = p.id", attributeSetID)
			joinExpr2 := facet.In(fmt.Sprintf("pa_%d.attribute_value_id", attributeSetID), sqlbuilder.List(facetKeys))
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
	sb.Where(sb.EQ("c_parent.id", categoryID), "p.quantity > 0 and p.is_published = 1 and s.is_published = 1")
	sb.GroupBy("b.id", "b.name")
	sb.OrderBy("b.name")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Filter) getQueryFindAttributeFacetsByCategoryID(categoryID int64, selectedFilterKey string, applied model.FilterCombination) (string, []any) {
	facet := sqlbuilder.Select("p.id").From("product as p")
	for filterKey, facets := range applied {
		facetKeys := slices.Collect(maps.Keys(facets))
		if filterKey == selectedFilterKey {
			continue
		} else if filterKey == "S" {
			facet.Where(facet.In("p.supplier_id", sqlbuilder.List(facetKeys)))
		} else if filterKey == "B" {
			facet.Where(facet.In("p.brand_id", sqlbuilder.List(facetKeys)))
		} else if filterKey == "PF" && len(facetKeys) == 1 {
			facet.Where(facet.GTE("p.web_actual", facetKeys[0]))
		} else if filterKey == "PT" && len(facetKeys) == 1 {
			facet.Where(facet.LTE("p.web_actual", facetKeys[0]))
		} else if attributeSetID, err := strconv.Atoi(filterKey); err == nil {
			joinTable := fmt.Sprintf("product_attribute as pa_%d", attributeSetID)
			joinExpr1 := fmt.Sprintf("pa_%d.product_id = p.id", attributeSetID)
			joinExpr2 := facet.In(fmt.Sprintf("pa_%d.attribute_value_id", attributeSetID), sqlbuilder.List(facetKeys))
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
		sb.EQ("p.category_id", categoryID),
		sb.EQ("pa.attribute_set_id", selectedFilterKey),
		"p.quantity > 0 and p.is_published = 1 and s.is_published = 1",
	)
	sb.GroupBy("atv.id", "atv.name")
	sb.OrderBy("atv.position")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Filter) getQueryFindPriceLimitByCategoryID(selectedCategoryID int64) (string, []any) {
	sb := sqlbuilder.Select(
		"coalesce(cast(min(p.web_actual) as int), 0) as price_from",
		"coalesce(round(max(p.web_actual + 0.4)), 0) as price_to",
	)
	sb.From("product as p")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "supplier as s", "s.id = p.supplier_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c", "c.id = p.category_id")
	sb.JoinWithOption(sqlbuilder.InnerJoin, "category as c_parent", "c.mp_path like c_parent.mp_path || '%'")
	sb.Where(sb.EQ("c_parent.id", selectedCategoryID), "p.quantity > 0 and p.is_published = 1 and s.is_published = 1")

	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *Filter) FindPriceLimitByCategoryID(ctx context.Context, categoryID int64) (model.FilterPrice, error) {
	const op = "sqlitedb.Filter.FindPriceLimitByCategoryID"
	query, args := st.getQueryFindPriceLimitByCategoryID(categoryID)
	price, _, err := runGet[model.FilterPrice](ctx, st.db, query, args)
	return price, utils.WrapIfErr(op, err)
}
