package model

import "gostart-crm/internal/app/storage"

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

func (r W2GridDataRequest) ToFindManyParams() storage.FindManyParams {
	p := storage.FindManyParams{
		Limit:  r.Limit,
		Offset: r.Offset,
	}

	if r.SeachLogic == "AND" {
		p.LogicAnd = true
	}

	if len(r.Search) > 0 {
		p.Filters = make([]storage.QueryFilter, len(r.Search))
		for i, s := range r.Search {
			p.Filters[i] = storage.QueryFilter{
				Field:    s.Field,
				Operator: s.Operator,
				Value:    s.Value,
			}
		}
	}

	if len(r.Sort) > 0 {
		p.Sorters = make([]storage.QuerySorter, len(r.Sort))
		for i, s := range r.Sort {
			if s.Direction == "desc" {
				p.Sorters[i] = storage.QuerySorter{
					Field: s.Field,
					Desc:  true,
				}
			} else {
				p.Sorters[i] = storage.QuerySorter{
					Field: s.Field,
					Desc:  false,
				}
			}
		}
	}

	return p
}

type W2GridDataResponse[T any, V any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Records []T    `json:"records,omitempty"`
	Summary []V    `json:"summary,omitempty"`
	Total   int64  `json:"total,omitempty"`
}

type W2GridSaveRequest[T any] struct {
	Changes []T `json:"changes"`
}

type W2GridRemoveRequest struct {
	ID []int64 `json:"id"`
}

type W2GridReorderRequest struct {
	ID         int64 `json:"id"`
	MoveBefore int64 `json:"moveBefore"`
}

type W2DropdownRequest struct {
	Max    int    `json:"max"`
	Search string `json:"search"`
}

type W2Dropdown struct {
	ID   *int64  `json:"id" db:"id"`
	Text *string `json:"text" db:"text"`
}

type W2DropdownGenericResponse[T any] struct {
	Status  string `json:"status"`
	Records []T    `json:"records"`
}

type W2DropdownBasicResponse = W2DropdownGenericResponse[W2Dropdown]

type W2FormRequest[T any] struct {
	Cmd    string `json:"cmd"`
	Name   string `json:"name"`
	RecID  int64  `json:"recid"`
	Record T      `json:"record"`
}

type W2FormResponse[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	RecID   int64  `json:"recid"`
	Record  *T     `json:"record,omitempty"`
}
