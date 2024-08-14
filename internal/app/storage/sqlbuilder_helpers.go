package storage

import (
	"fmt"

	"github.com/huandu/go-sqlbuilder"
)

type QueryFilter struct {
	Field    string
	Operator string
	Value    any
}

type QuerySorter struct {
	Field string
	Desc  bool
}

type FindManyParams struct {
	Limit    int
	Offset   int
	LogicAnd bool
	Filters  []QueryFilter
	Sorters  []QuerySorter
}

func ApplyLimitOffset(sb *sqlbuilder.SelectBuilder, limit int, offset int) {
	if limit != 0 {
		sb.Limit(limit)
	}
	if offset != 0 {
		sb.Offset(offset)
	}
}

func ApplyFilters(sb *sqlbuilder.SelectBuilder, qf []QueryFilter, logicAnd bool, allowed map[string]string) {
	c := make([]string, 0, len(qf))

	for _, f := range qf {
		field, ok := allowed[f.Field]
		if !ok {
			continue
		}

		switch f.Operator {
		case "=", "is":
			c = append(c, sb.EQ(field, f.Value))
		case ">":
			c = append(c, sb.GT(field, f.Value))
		case "<", "less":
			c = append(c, sb.LT(field, f.Value))
		case ">=", "more":
			c = append(c, sb.GTE(field, f.Value))
		case "<=":
			c = append(c, sb.LTE(field, f.Value))
		case "begins":
			c = append(c, sb.Like(field, fmt.Sprintf("%v%%", f.Value)))
		case "contains":
			c = append(c, sb.Like(field, fmt.Sprintf("%%%v%%", f.Value)))
		case "ends":
			c = append(c, sb.Like(field, fmt.Sprintf("%%%v", f.Value)))
		case "between":
			if values, ok := f.Value.([]any); ok && len(values) == 2 {
				c = append(c, sb.Between(field, values[0], values[1]))
			}
		case "in":
			if values, ok := f.Value.([]any); ok {
				ids := make([]any, len(values))
				for i := range values {
					if value, ok := values[i].(map[string]any); ok {
						ids[i] = value["id"]
					}
				}
				c = append(c, sb.In(field, ids...))
			}
		case "not in":
			if values, ok := f.Value.([]any); ok {
				ids := make([]any, len(values))
				for i := range values {
					if value, ok := values[i].(map[string]any); ok {
						ids[i] = value["id"]
					}
				}
				c = append(c, sb.NotIn(field, ids...))
			}
		}
	}

	if len(c) > 0 {
		if logicAnd {
			sb.Where(sb.And(c...))
		} else {
			sb.Where(sb.Or(c...))
		}
	}
}

func ApplySorters(sb *sqlbuilder.SelectBuilder, qs []QuerySorter, allowed map[string]string) {
	for _, s := range qs {
		safeField, ok := allowed[s.Field]
		if !ok {
			continue
		}

		if s.Desc {
			sb.OrderBy(fmt.Sprintf("%s DESC", safeField))
		} else {
			sb.OrderBy(fmt.Sprintf("%s ASC", safeField))
		}
	}
}

func ApplyUpdateSetPartial(ub *sqlbuilder.UpdateBuilder, partialMap map[string]struct{}, partialKey string, field string, value any) {
	if _, ok := partialMap[partialKey]; !ok {
		return
	}

	if strValue, ok := value.(string); ok && strValue == "" {
		ub.SetMore(ub.EQ(field, nil))
	} else {
		ub.SetMore(ub.EQ(field, value))
	}
}
