package sqlitedb

import (
	"fmt"
	"strings"

	"gostart-crm/internal/app/model"

	"github.com/huandu/go-sqlbuilder"
)

func QuoteIdentifier(name string) string {
	return `"` + strings.Replace(name, `"`, `""`, -1) + `"`
}

func ApplyQueryLimitOffset(sb *sqlbuilder.SelectBuilder, limit int, offset int) {
	if limit != 0 {
		sb.Limit(limit)
	}
	if offset != 0 {
		sb.Offset(offset)
	}
}

func ApplyQueryFilters(sb *sqlbuilder.SelectBuilder, qf [][]model.QueryFilter) {
	for _, g := range qf {
		og := make([]string, len(g))

		for i, w := range g {
			safeField := QuoteIdentifier(w.Field)
			switch w.Operator {
			case "=", "is":
				og[i] = sb.EQ(safeField, w.Value)
			case ">":
				og[i] = sb.GT(safeField, w.Value)
			case "<":
				og[i] = sb.LT(safeField, w.Value)
			case ">=":
				og[i] = sb.GTE(safeField, w.Value)
			case "<=":
				og[i] = sb.LTE(safeField, w.Value)
			case "begins":
				og[i] = sb.Like(safeField, fmt.Sprintf("%v%%", w.Value))
			case "contains":
				og[i] = sb.Like(safeField, fmt.Sprintf("%%%v%%", w.Value))
			case "ends":
				og[i] = sb.Like(safeField, fmt.Sprintf("%%%v", w.Value))
			case "between":
				if values, ok := w.Value.([]any); ok && len(values) == 2 {
					og[i] = sb.Between(safeField, values[0], values[1])
				}
			}
		}

		sb.Where(sb.Or(og...))
	}
}

func ApplyQuerySorters(sb *sqlbuilder.SelectBuilder, qs []model.QuerySorter) {
	for _, s := range qs {
		safeField := QuoteIdentifier(s.Field)
		if s.Desc {
			sb.OrderBy(fmt.Sprintf("%s DESC", safeField))
		} else {
			sb.OrderBy(fmt.Sprintf("%s ASC", safeField))
		}
	}
}
