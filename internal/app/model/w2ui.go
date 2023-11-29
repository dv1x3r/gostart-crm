package model

type W2BaseResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type W2GridSearch struct {
	Field    string `json:"field"`
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Value    any    `json:"value"`
}

type W2GridSort struct {
	Field     string `json:"field"`
	Direction string `json:"direction"`
}

type W2GridDataRequest struct {
	Limit      int            `json:"limit"`
	Offset     int            `json:"offset"`
	SeachLogic string         `json:"searchLogic"`
	Search     []W2GridSearch `json:"search"`
	Sort       []W2GridSort   `json:"sort"`
}

type W2GridDataResponse[T any, V any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Records []T    `json:"records,omitempty"`
	Summary []V    `json:"summary,omitempty"`
	Total   int64  `json:"total,omitempty"`
}

type W2GridPatchRequest[T any] struct {
	Changes []T `json:"changes"`
}

type W2GridDeleteRequest struct {
	ID []int64 `json:"id"`
}

type W2FormRequest[T any] struct {
	Cmd    string `json:"cmd"`
	Name   string `json:"name"`
	RecID  int64  `json:"recid"`
	Record T      `json:"record"`
}

type W2FormResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Record  T      `json:"record"`
}

func (r W2GridDataRequest) ToQueryListParams() QueryListParams {
	q := QueryListParams{
		Limit:  r.Limit,
		Offset: r.Offset,
	}

	if len(r.Search) > 0 {
		q.Where = [][]QueryWhere{{}}
	}

	for _, s := range r.Search {
		q.Where[0] = append(q.Where[0], QueryWhere{
			Field:    s.Field,
			Operator: s.Operator,
			Value:    s.Value,
		})
	}

	for _, s := range r.Sort {
		if s.Direction == "desc" {
			q.OrderBy = append(q.OrderBy, QueryOrderBy{
				Field: s.Field,
				Desc:  true,
			})
		} else {
			q.OrderBy = append(q.OrderBy, QueryOrderBy{
				Field: s.Field,
				Desc:  false,
			})
		}
	}

	return q
}
