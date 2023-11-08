package model

type WGrid[T any] struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Total   int64  `json:"total,omitempty"`
	Records []T    `json:"records"`
	Summary any    `json:"summary,omitempty"`
}
