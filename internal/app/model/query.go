package model

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
	Limit   int
	Offset  int
	Filters [][]QueryFilter
	Sorters []QuerySorter
}
