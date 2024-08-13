package service

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
)

type TodoStorager interface {
	FindMany(context.Context, storage.FindManyParams) ([]model.TodoDTO, error)
	GetOneByID(context.Context, int64) (model.TodoDTO, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	PatchManyByID(context.Context, []model.TodoDTO) (int64, error)
	UpdateByID(context.Context, model.TodoDTO) (int64, error)
	Insert(context.Context, model.TodoDTO) (int64, error)
}

type Todo struct {
	st TodoStorager
}

func NewTodo(st TodoStorager) *Todo {
	return &Todo{st: st}
}

func (ts *Todo) GetTodoList(ctx context.Context) ([]model.TodoDTO, error) {
	return ts.st.FindMany(ctx, storage.FindManyParams{})
}

func (ts *Todo) GetTodoW2Grid(ctx context.Context, req model.W2GridDataRequest) (model.TodoW2GridResponse, error) {
	rows, err := ts.st.FindMany(ctx, req.ToFindManyParams())
	if err != nil {
		return model.TodoW2GridResponse{Status: "error"}, err
	}

	todos := make([]model.TodoDTO, len(rows))
	var count int64
	for i, row := range rows {
		// count = row.Count
		todos[i] = row //.TodoDTO
	}

	return model.TodoW2GridResponse{Status: "success", Records: todos, Total: count}, nil
}

func (ts *Todo) DeleteTodoW2Grid(ctx context.Context, req model.W2GridRemoveRequest) (model.W2BaseResponse, error) {
	if _, err := ts.st.DeleteManyByID(ctx, req.ID); err != nil {
		return model.W2BaseResponse{Status: "error"}, err
	}
	return model.W2BaseResponse{Status: "success"}, nil
}

func (ts *Todo) PatchTodoW2Action(ctx context.Context, req model.TodoW2SaveRequest) (model.W2BaseResponse, error) {
	if _, err := ts.st.PatchManyByID(ctx, req.Changes); err != nil {
		return model.W2BaseResponse{Status: "error"}, err
	}
	return model.W2BaseResponse{Status: "success"}, nil
}
