package service

import (
	"context"

	"w2go/internal/app/model"
)

type TodoStorager interface {
	FindMany(context.Context, model.FindManyParams) ([]model.TodoFromDB, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	// Insert(model.TodoDTO) (int64, error)
	// Update(int64, model.TodoDTO) error
	// Patch(int64, model.TodoPatchDTO) error
}

type Todo struct {
	st TodoStorager
}

func NewTodo(st TodoStorager) *Todo {
	return &Todo{st: st}
}

func (s *Todo) GetTodoW2Grid(ctx context.Context, req model.W2GridDataRequest) (model.TodoW2GridResponse, error) {
	rows, err := s.st.FindMany(ctx, req.ToFindManyParams())
	if err != nil {
		return model.TodoW2GridResponse{Status: "error"}, err
	}

	todos := make([]model.TodoDTO, len(rows))
	var count int64
	for i, row := range rows {
		count = row.Count
		todos[i] = row.TodoDTO
	}

	return model.TodoW2GridResponse{Status: "success", Records: todos, Total: count}, nil
}

func (s *Todo) DeleteTodoW2Grid(ctx context.Context, req model.W2GridDeleteRequest) (model.W2BaseResponse, error) {
	if _, err := s.st.DeleteManyByID(ctx, req.ID); err != nil {
		return model.W2BaseResponse{Status: "error"}, err
	}
	return model.W2BaseResponse{Status: "success"}, nil
}

func (ts *Todo) PatchTodoW2Action(ctx context.Context, req model.TodoW2PatchRequest) (model.W2BaseResponse, error) {
	return model.W2BaseResponse{Status: "error", Message: "not implemented"}, nil
}

func (ts *Todo) GetTodoW2Form(ctx context.Context, req model.TodoW2FormRequest) (model.TodoW2FormResponse, error) {
	// for _, row := range ts.data {
	// 	if row.ID == req.RecID {
	// 		return model.TodoW2FormResponse{Status: "success", Record: row}, nil
	// 	}
	// }
	return model.TodoW2FormResponse{Status: "error", Message: "not found"}, nil
}

func (ts *Todo) UpsertTodoW2Form(ctx context.Context, req model.TodoW2FormRequest) (model.W2BaseResponse, error) {
	// if req.RecID == 0 {
	// 	ts.data = append(ts.data, req.Record)
	// } else {
	// 	for i, row := range ts.data {
	// 		if row.ID == req.RecID {
	// 			ts.data[i] = req.Record
	// 			break
	// 		}
	// 	}
	// }
	return model.W2BaseResponse{Status: "success"}, nil
}
