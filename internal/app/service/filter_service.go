package service

import (
	"context"
	"strconv"
	"strings"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/utils"
)

type FilterStorager interface {
	FindFacetsByCategoryID(context.Context, int64, model.FilterCombination) ([]model.FilterFacet, error)
}

type Filter struct {
	filterStorage FilterStorager
}

func NewFilter(filterStorage FilterStorager) *Filter {
	return &Filter{filterStorage: filterStorage}
}

func (sc *Filter) ParseAppliedFilters(query string) model.FilterCombination {
	applied := make(model.FilterCombination)

	for _, filter := range strings.Split(query, "_") {
		split := strings.Split(filter, "-")
		if len(split) != 2 {
			continue
		}

		key := split[0]
		value, _ := strconv.Atoi(split[1])
		if applied[key] == nil {
			applied[key] = make(map[int64]struct{})
		}

		applied[key][int64(value)] = struct{}{}
	}

	return applied
}

func (sc *Filter) GetFacetsByCategoryID(ctx context.Context, categoryID int64, applied model.FilterCombination) ([]model.FilterFacet, error) {
	const op = "service.Filter.GetFacetsByCategoryID"
	facets, err := sc.filterStorage.FindFacetsByCategoryID(ctx, categoryID, applied)
	return facets, utils.WrapIfErr(op, err)
}
