package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrCategoryExists          = errors.New("service.Category: already exists")
	ErrCategoryNotFound        = errors.New("service.Category: not found")
	ErrCategoryHasRelated      = errors.New("service.Category: has related objects")
	ErrCategoryInvalidReorder  = errors.New("service.Category: invalid reorder parameters")
	ErrCategoryInvalidCircular = errors.New("service.Category: invalid circular tree")
)

type CategoryStorager interface {
	GetTree(context.Context) (model.Category, error)
	GetDropdown(context.Context, string, int, bool) ([]model.CategoryDropdown, error)
	GetByID(context.Context, int64) (model.Category, bool, error)
	FindManyByParentID(context.Context, int64, storage.FindManyParams) ([]model.Category, int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	UpsertMany(context.Context, []model.Category) (int64, error)
	UpdatePositions(context.Context, []int64) (int64, error)
}

type Category struct {
	categoryStorage CategoryStorager
}

func NewCategory(categoryStorage CategoryStorager) *Category {
	return &Category{categoryStorage: categoryStorage}
}

func (sc *Category) FetchTree(ctx context.Context) (model.Category, error) {
	const op = "service.Category.FetchTree"
	tree, err := sc.categoryStorage.GetTree(ctx)
	return tree, utils.WrapIfErr(op, err)
}

func (sc *Category) FetchDropdown(ctx context.Context, search string, max int, leafsOnly bool) ([]model.CategoryDropdown, error) {
	const op = "service.Category.FetchDropdown"
	rows, err := sc.categoryStorage.GetDropdown(ctx, search, max, leafsOnly)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *Category) FetchChildren(ctx context.Context, parentID int64, params storage.FindManyParams) ([]model.Category, int64, error) {
	const op = "service.Category.FetchChildren"
	rows, count, err := sc.categoryStorage.FindManyByParentID(ctx, parentID, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *Category) SavePartials(ctx context.Context, dtos []model.Category) (int64, error) {
	const op = "service.Category.SavePartials"

	for i, dto := range dtos {
		if dto.ParentID != nil && *dto.ParentID == 0 {
			dtos[i].ParentID = nil // set root category
		}
	}

	affected, err := sc.categoryStorage.UpsertMany(ctx, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return affected, utils.WrapIfErr(op, errors.Join(ErrCategoryExists, err))
	} else if errors.Is(err, storage.ErrInvalidTreeCircular) {
		return affected, utils.WrapIfErr(op, errors.Join(ErrCategoryInvalidCircular, err))
	} else if err != nil {
		return affected, utils.WrapIfErr(op, err)
	}

	return affected, nil
}

func (sc *Category) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.Category.DeleteListID"
	affected, err := sc.categoryStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrCategoryHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *Category) Reorder(ctx context.Context, objectID int64, moveBefore int64) error {
	const op = "service.Category.Reorder"

	object, ok, err := sc.categoryStorage.GetByID(ctx, objectID)
	if err != nil {
		return utils.WrapIfErr(op, err)
	} else if !ok {
		return utils.WrapIfErr(op, ErrCategoryNotFound)
	}

	var parentID int64
	if object.ParentID != nil {
		parentID = *object.ParentID
	}

	group, _, err := sc.categoryStorage.FindManyByParentID(ctx, parentID, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	ids := make([]int64, len(group))
	for i, value := range group {
		ids[i] = value.ID
	}

	if ok := utils.ReorderBefore(ids, objectID, moveBefore); !ok {
		return utils.WrapIfErr(op, ErrCategoryInvalidReorder)
	}

	if _, err := sc.categoryStorage.UpdatePositions(ctx, ids); err != nil {
		return utils.WrapIfErr(op, err)
	}

	return nil
}
