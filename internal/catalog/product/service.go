package product

import (
	"context"
	"database/sql"
	"errors"

	"gostart-crm/internal/lib/wrap"

	"github.com/dv1x3r/w2go/w2"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Fetch(ctx context.Context, r w2.GridDataRequest, categoryID int) ([]Product, int, error) {
	const op = "product.Service.Fetch"
	rows, count, err := s.repo.FindManyByCategoryID(ctx, r, categoryID)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) FetchList(ctx context.Context, r ListDataRequest, categoryID int) ([]Product, error) {
	const op = "product.Service.FetchList"
	rows, err := s.repo.FindAvailableByCategoryID(ctx, r, categoryID)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) FetchByID(ctx context.Context, id int) (Product, error) {
	const op = "product.Service.FetchByID"
	product, err := s.repo.GetByID(ctx, id)
	if errors.Is(err, sql.ErrNoRows) {
		return product, wrap.IfErr(op, ErrProductNotFound)
	}
	return product, wrap.IfErr(op, err)
}

func (s *Service) SaveForm(ctx context.Context, dto Product) (int, error) {
	const op = "product.Service.SaveForm"

	resultID, err := s.repo.UpsertOne(ctx, dto)
	if s.repo.Store().IsErrConstraintUnique(err) {
		return 0, wrap.IfErr(op, ErrProductExists)
	}

	return resultID, wrap.IfErr(op, err)
}

func (s *Service) SaveChanges(ctx context.Context, changes []Product) error {
	const op = "product.Service.SaveChanges"
	err := s.repo.UpdateMany(ctx, changes)
	return wrap.IfErr(op, err)
}

func (s *Service) DeleteListID(ctx context.Context, ids []int) error {
	const op = "product.Service.DeleteListID"

	if len(ids) > 50 {
		return wrap.IfErr(op, ErrProductDeleteProtect)
	}

	if err := s.repo.DeleteManyByID(ctx, ids); err != nil {
		return wrap.IfErr(op, err)
	}

	return nil
}

func (s *Service) FetchAttributes(ctx context.Context, productID int, r w2.GridDataRequest) ([]ProductAttribute, int, error) {
	const op = "product.Service.FetchAttributes"
	rows, count, err := s.repo.FindManyAttributesByProductID(ctx, productID, r)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) SaveAttributeChanges(ctx context.Context, productID int, changes []ProductAttribute) error {
	const op = "product.Service.SaveAttributeChanges"
	err := s.repo.UpsertManyAttributes(ctx, productID, changes)
	if s.repo.Store().IsErrConstraintForeignKey(err) {
		return wrap.IfErr(op, ErrProductAttributeNotFound)
	}
	return wrap.IfErr(op, err)
}
