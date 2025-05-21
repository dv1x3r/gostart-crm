package category

import (
	"context"
	"database/sql"
	"errors"

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

func (s *Service) FetchTree(ctx context.Context, useCache bool) (Category, error) {
	const op = "category.Service.FetchTree"
	tree, err := s.repo.GetTree(ctx, useCache)
	return tree, wrap.IfErr(op, err)
}

func (s *Service) GetByURL(ctx context.Context, searchURL string) (Category, error) {
	const op = "category.Service.GetByURL"

	root, err := s.FetchTree(ctx, true)
	if err != nil {
		return Category{}, wrap.IfErr(op, err)
	}

	category, found := s.getByURLRecursive(root, searchURL)
	if !found {
		return category, wrap.IfErr(op, ErrCategoryNotFound)
	}

	return category, nil
}

func (s *Service) getByURLRecursive(category Category, searchURL string) (Category, bool) {
	if category.CategoryURL == searchURL {
		return category, true
	}

	for _, child := range category.Children {
		if target, ok := s.getByURLRecursive(child, searchURL); ok {
			return target, true
		}
	}

	return category, false
}

func (s *Service) FetchDropdown(ctx context.Context, r w2.DropdownRequest, leafsOnly bool) ([]w2.DropdownValue, error) {
	const op = "category.Service.FetchDropdown"
	rows, err := s.repo.GetDropdown(ctx, r, leafsOnly)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) FetchChildren(ctx context.Context, parentID int, r w2.GridDataRequest) ([]Category, int, error) {
	const op = "category.Service.FetchChildren"
	rows, count, err := s.repo.FindManyByParentID(ctx, parentID, r)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) SaveChanges(ctx context.Context, changes []Category) error {
	const op = "category.Service.SaveChanges"
	err := s.repo.UpsertMany(ctx, changes)
	if s.repo.Store().IsErrConstraintUnique(err) {
		return wrap.IfErr(op, ErrCategoryExists)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) DeleteListID(ctx context.Context, ids []int) error {
	const op = "category.Service.DeleteListID"
	err := s.repo.DeleteManyByID(ctx, ids)
	if s.repo.Store().IsErrConstraintTrigger(err) {
		return wrap.IfErr(op, ErrCategoryHasRelated)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) Reorder(ctx context.Context, r w2.GridReorderRequest) error {
	const op = "category.Service.Reorder"

	parentID, err := s.repo.GetParentIDByCategoryID(ctx, r.RecID)
	if errors.Is(err, sql.ErrNoRows) {
		return wrap.IfErr(op, ErrCategoryNotFound)
	} else if err != nil {
		return wrap.IfErr(op, err)
	}

	categories, _, err := s.repo.FindManyByParentID(ctx, parentID, w2.GridDataRequest{})
	if err != nil {
		return wrap.IfErr(op, err)
	}

	ids := make([]int, len(categories))
	for i, value := range categories {
		ids[i] = value.ID
	}

	if err := w2sort.ReorderArray(ids, r); err != nil {
		return wrap.IfErr(op, ErrCategoryInvalidReorder)
	}

	return wrap.IfErr(op, s.repo.UpdatePositions(ctx, ids))
}
