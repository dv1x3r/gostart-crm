package model

type FilterCombination map[string]map[int64]struct{}

type FilterFacetValue struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Products   int64  `db:"products"`
	IsSelected bool
}

type FilterFacet struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Values   []FilterFacetValue
	Selected int
}

type FilterPrice struct {
	From int64 `db:"price_from"`
	To   int64 `db:"price_to"`
}

type Filter struct {
	Facets        []FilterFacet
	PriceLimit    FilterPrice
	PriceSelected FilterPrice
}
