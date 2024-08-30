package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrProductStatusExists         = errors.New("service.ProductStatus: already exists")
	ErrProductStatusHasRelated     = errors.New("service.ProductStatus: has related objects")
	ErrProductStatusInvalidReorder = errors.New("service.ProductStatus: invalid reorder parameters")
)

type ProductStatusStorager interface {
	FindMany(context.Context, storage.FindManyParams) ([]model.ProductStatus, int64, error)
	GetDropdown(context.Context, string, int) ([]model.ProductStatusDropdown, error)
	UpsertMany(context.Context, []model.ProductStatus) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	UpdatePositions(context.Context, []int64) (int64, error)
}

type ProductStatus struct {
	productStatusStorage ProductStatusStorager
}

func NewProductStatus(productStatusStorage ProductStatusStorager) *ProductStatus {
	return &ProductStatus{productStatusStorage: productStatusStorage}
}

func (sc *ProductStatus) Fetch(ctx context.Context, params storage.FindManyParams) ([]model.ProductStatus, int64, error) {
	const op = "service.ProductStatus.Fetch"
	rows, count, err := sc.productStatusStorage.FindMany(ctx, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *ProductStatus) FetchDropdown(ctx context.Context, search string, max int) ([]model.ProductStatusDropdown, error) {
	const op = "service.ProductStatus.FetchDropdown"
	rows, err := sc.productStatusStorage.GetDropdown(ctx, search, max)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *ProductStatus) SavePartials(ctx context.Context, dtos []model.ProductStatus) (int64, error) {
	const op = "service.ProductStatus.SavePartials"
	affected, err := sc.productStatusStorage.UpsertMany(ctx, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrProductStatusExists, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *ProductStatus) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.ProductStatus.DeleteListID"
	affected, err := sc.productStatusStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrProductStatusHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *ProductStatus) Reorder(ctx context.Context, objectID int64, moveBefore int64) error {
	const op = "service.ProductStatus.Reorder"

	group, _, err := sc.productStatusStorage.FindMany(ctx, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	ids := make([]int64, len(group))
	for i, value := range group {
		ids[i] = value.ID
	}

	if ok := utils.ReorderBefore(ids, objectID, moveBefore); !ok {
		return utils.WrapIfErr(op, ErrProductStatusInvalidReorder)
	}

	if _, err := sc.productStatusStorage.UpdatePositions(ctx, ids); err != nil {
		return utils.WrapIfErr(op, err)
	}

	return nil
}
