package model

type QueryFilter struct {
	Field    string
	Operator string
	Value    any
}

type QuerySort struct {
	Field string
	Desc  bool
}
