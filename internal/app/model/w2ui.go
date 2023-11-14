package model

type W2GridSearch struct {
	Field    string `query:"field"`
	Type     string `query:"type"`
	Operator string `query:"operator"`
	Value    any    `query:"value"`
}

type W2GridSort struct {
	Field     string `query:"field"`
	Direction string `query:"direction"`
}

type W2GridDataRequest struct {
	Limit      int64        `query:"limit"`
	Offset     int64        `query:"offset"`
	SeachLogic string       `query:"searchLogic"`
	Search     W2GridSearch `query:"search"`
	Sort       W2GridSort   `query:"sort"`
}

type W2GridDataResponse[T any, V any] struct {
	Status  string `json:"status"`
	Records []T    `json:"records"`
	Summary []V    `json:"summary,omitempty"`
	Message string `json:"message,omitempty"`
	Total   int64  `json:"total,omitempty"`
}

type W2GridActionRequest[T any] struct {
	Action  string `json:"action"`
	RecID   []T    `json:"recid,omitempty"`   // Action: delete
	Changes []T    `json:"changes,omitempty"` // Action: save
}

type W2GridActionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type W2FormSubmit[T any, V any] struct {
	Cmd    string `json:"cmd"`
	Name   string `json:"name"`
	RecID  V      `json:"recid"`
	Record T      `json:"record"`
}

type W2FormResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
