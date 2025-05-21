package brand

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

func (s *Service) Fetch(ctx context.Context, r w2.GridDataRequest) ([]Brand, int, error) {
	const op = "brand.Service.Fetch"
	rows, count, err := s.repo.FindMany(ctx, r)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) FetchDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "brand.Service.FetchDropdown"
	rows, err := s.repo.GetDropdown(ctx, r)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) SaveChanges(ctx context.Context, changes []Brand) error {
	const op = "brand.Service.SaveChanges"
	err := s.repo.UpsertMany(ctx, changes)
	if s.repo.Store().IsErrConstraintUnique(err) {
		return wrap.IfErr(op, ErrBrandExists)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) DeleteListID(ctx context.Context, ids []int) error {
	const op = "brand.Service.DeleteListID"
	err := s.repo.DeleteManyByID(ctx, ids)
	if s.repo.Store().IsErrConstraintTrigger(err) {
		return wrap.IfErr(op, ErrBrandHasRelated)
	}
	return wrap.IfErr(op, err)
}
