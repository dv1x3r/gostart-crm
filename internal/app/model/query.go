package model

type QueryWhere struct {
	Field    string
	Operator string
	Value    any
}

type QueryOrderBy struct {
	Field string
	Desc  bool
}

type QueryListParams struct {
	Limit   int
	Offset  int
	Where   [][]QueryWhere
	OrderBy []QueryOrderBy
}
