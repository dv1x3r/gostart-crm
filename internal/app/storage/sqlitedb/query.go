package sqlitedb

import (
	"fmt"
	"strings"

	"w2go/internal/app/model"

	"github.com/huandu/go-sqlbuilder"
)

func QuoteIdentifier(name string) string {
	end := strings.IndexRune(name, 0)
	if end > -1 {
		name = name[:end]
	}
	return `"` + strings.Replace(name, `"`, `""`, -1) + `"`
}

func ApplyQueryFilter(sb *sqlbuilder.SelectBuilder, qf [][]model.QueryFilter) {
	for _, g := range qf {
		og := []string{}

		for _, w := range g {
			safeField := QuoteIdentifier(w.Field)
			switch w.Operator {
			case "=", "is":
				og = append(og, sb.EQ(safeField, w.Value))
			case ">":
				og = append(og, sb.GT(safeField, w.Value))
			case "<":
				og = append(og, sb.LT(safeField, w.Value))
			case ">=":
				og = append(og, sb.GTE(safeField, w.Value))
			case "<=":
				og = append(og, sb.LTE(safeField, w.Value))
			case "begins":
				og = append(og, sb.Like(safeField, fmt.Sprintf("%v%%", w.Value)))
			case "contains":
				og = append(og, sb.Like(safeField, fmt.Sprintf("%%%v%%", w.Value)))
			case "ends":
				og = append(og, sb.Like(safeField, fmt.Sprintf("%%%v", w.Value)))
			case "between":
				if values, ok := w.Value.([]any); ok && len(values) == 2 {
					fmt.Println("YES")
					og = append(og, sb.Between(safeField, values[0], values[1]))
				}
			}
		}

		sb.Where(sb.Or(og...))
	}
}

func applyQueryOrderBy(sb *sqlbuilder.SelectBuilder, qo []model.QuerySort) {
	for _, s := range qo {
		safeField := QuoteIdentifier(s.Field)
		if s.Desc {
			sb = sb.OrderBy(fmt.Sprintf("%s DESC", safeField))
		} else {
			sb = sb.OrderBy(fmt.Sprintf("%s ASC", safeField))
		}
	}
}
