package model

type FilterCombination map[string]map[int64]struct{}

type FilterFacet struct {
	ID         int64  `db:"id"`
	Name       string `db:"name"`
	Products   int64  `db:"products"`
	IsSelected bool
}

type Filter struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Facets   []FilterFacet
	Selected int
}
