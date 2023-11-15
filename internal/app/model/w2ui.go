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
	Limit      int64          `json:"limit"`
	Offset     int64          `json:"offset"`
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

type W2FormSubmit[T any] struct {
	Cmd    string `json:"cmd"`
	Name   string `json:"name"`
	RecID  int64  `json:"recid"`
	Record T      `json:"record"`
}
