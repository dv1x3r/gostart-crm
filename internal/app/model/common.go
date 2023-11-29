package model

type QuerySearch struct {
	Field    string
	Type     string
	Operator string
	Value    any
}

type QuerySort struct {
	Field     string
	Direction string
}

type QueryList struct {
	Limit  int64
	Offset int64
	Search []QuerySearch
	Sort   []QuerySort
}
