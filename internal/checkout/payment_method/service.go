package paymentmethod

import (
	"context"

	"gostart-crm/internal/lib/wrap"

	"github.com/dv1x3r/w2go/w2"
	"github.com/dv1x3r/w2go/w2sort"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Fetch(ctx context.Context, r w2.GridDataRequest) ([]PaymentMethod, int, error) {
	const op = "paymentmethod.Service.Fetch"
	rows, count, err := s.repo.FindMany(ctx, r)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) FetchAll(ctx context.Context) ([]PaymentMethod, error) {
	const op = "paymentmethod.Service.FetchAll"
	rows, err := s.repo.FindAll(ctx)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) FetchDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "paymentmethod.Service.FetchDropdown"
	rows, err := s.repo.GetDropdown(ctx, r)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) SaveChanges(ctx context.Context, changes []PaymentMethod) error {
	const op = "paymentmethod.Service.SaveChanges"
	err := s.repo.UpsertMany(ctx, changes)
	if s.repo.Store().IsErrConstraintUnique(err) {
		return wrap.IfErr(op, ErrPaymentMethodExists)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) DeleteListID(ctx context.Context, ids []int) error {
	const op = "paymentmethod.Service.DeleteListID"
	err := s.repo.DeleteManyByID(ctx, ids)
	if s.repo.Store().IsErrConstraintTrigger(err) {
		return wrap.IfErr(op, ErrPaymentMethodHasRelated)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) Reorder(ctx context.Context, r w2.GridReorderRequest) error {
	const op = "paymentmethod.Service.Reorder"

	paymentMethods, _, err := s.repo.FindMany(ctx, w2.GridDataRequest{})
	if err != nil {
		return wrap.IfErr(op, err)
	}

	ids := make([]int, len(paymentMethods))
	for i, value := range paymentMethods {
		ids[i] = value.ID
	}

	if err := w2sort.ReorderArray(ids, r); err != nil {
		return wrap.IfErr(op, ErrPaymentMethodInvalidReorder)
	}

	return wrap.IfErr(op, s.repo.UpdatePositions(ctx, ids))
}
