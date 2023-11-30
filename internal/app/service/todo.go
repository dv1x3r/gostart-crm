package service

import (
	"context"

	"w2go/internal/app/model"
)

type TodoStorager interface {
	FindMany(context.Context, model.FindManyParams) ([]model.TodoFromDB, error)
	GetOneByID(context.Context, int64) (model.TodoDTO, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	PatchManyByID(context.Context, []model.TodoPartialDTO) (int64, error)
	UpdateByID(context.Context, model.TodoDTO) (int64, error)
	Insert(context.Context, model.TodoDTO) (int64, error)
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

func (s *Todo) PatchTodoW2Action(ctx context.Context, req model.TodoW2PatchRequest) (model.W2BaseResponse, error) {
	if _, err := s.st.PatchManyByID(ctx, req.Changes); err != nil {
		return model.W2BaseResponse{Status: "error"}, err
	}
	return model.W2BaseResponse{Status: "success"}, nil
}

func (s *Todo) GetTodoW2Form(ctx context.Context, req model.TodoW2FormRequest) (model.TodoW2FormResponse, error) {
	todo, err := s.st.GetOneByID(ctx, req.RecID)
	if err != nil {
		return model.TodoW2FormResponse{Status: "error"}, err
	}
	return model.TodoW2FormResponse{Status: "success", Record: todo}, nil
}

func (s *Todo) UpsertTodoW2Form(ctx context.Context, req model.TodoW2FormRequest) (model.W2BaseResponse, error) {
	var err error

	if req.RecID == 0 {
		_, err = s.st.Insert(ctx, req.Record)
	} else {
		_, err = s.st.UpdateByID(ctx, req.Record)
	}

	if err != nil {
		return model.W2BaseResponse{Status: "error"}, err
	}

	return model.W2BaseResponse{Status: "success"}, nil
}
