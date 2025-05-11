package order

import (
	"context"

	"gostart-crm/internal/lib/wrap"

	"github.com/dv1x3r/w2go/w2"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Fetch(ctx context.Context, r w2.GridDataRequest) ([]Order, int, error) {
	const op = "order.Service.Fetch"
	rows, count, err := s.repo.FindMany(ctx, r)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) SaveForm(ctx context.Context, id int, dto Order) (int, error) {
	const op = "order.Service.SaveForm"
	err := s.repo.UpdateByID(ctx, id, dto)
	return id, wrap.IfErr(op, err)
}

func (s *Service) DeleteListID(ctx context.Context, ids []int) error {
	const op = "order.Service.DeleteListID"
	if len(ids) > 20 {
		return wrap.IfErr(op, ErrOrderDeleteProtect)
	}
	err := s.repo.DeleteManyByID(ctx, ids)
	return wrap.IfErr(op, err)
}
