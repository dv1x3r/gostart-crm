package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrSupplierExists         = errors.New("service.Supplier: already exists")
	ErrSupplierNotFound       = errors.New("service.Supplier: not found")
	ErrSupplierHasRelated     = errors.New("service.Supplier: has related objects")
	ErrSupplierInvalidReorder = errors.New("service.Supplier: invalid reorder parameters")
)

type SupplierStorager interface {
	FindMany(context.Context, storage.FindManyParams) ([]model.Supplier, int64, error)
	GetDropdown(context.Context, string, int) ([]model.W2Dropdown, error)
	GetByID(context.Context, int64) (model.Supplier, bool, error)
	UpsertMany(context.Context, []model.Supplier) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	UpdatePositions(context.Context, []int64) (int64, error)
}

type Supplier struct {
	supplierStorage SupplierStorager
}

func NewSupplier(supplierStorage SupplierStorager) *Supplier {
	return &Supplier{supplierStorage: supplierStorage}
}

func (sc *Supplier) Fetch(ctx context.Context, params storage.FindManyParams) ([]model.Supplier, int64, error) {
	const op = "service.Supplier.Fetch"
	rows, count, err := sc.supplierStorage.FindMany(ctx, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *Supplier) FetchDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "service.Supplier.FetchDropdown"
	rows, err := sc.supplierStorage.GetDropdown(ctx, search, max)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *Supplier) SavePartials(ctx context.Context, dtos []model.Supplier) (int64, error) {
	const op = "service.Supplier.SavePartials"
	affected, err := sc.supplierStorage.UpsertMany(ctx, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrSupplierExists, err))
	}

	return affected, utils.WrapIfErr(op, err)
}

func (sc *Supplier) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.Supplier.DeleteListID"
	affected, err := sc.supplierStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrSupplierHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *Supplier) Reorder(ctx context.Context, objectID int64, moveBefore int64) error {
	const op = "service.Supplier.Reorder"

	group, _, err := sc.supplierStorage.FindMany(ctx, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	ids := make([]int64, len(group))
	for i, value := range group {
		ids[i] = value.ID
	}

	if ok := utils.ReorderBefore(ids, objectID, moveBefore); !ok {
		return utils.WrapIfErr(op, ErrSupplierInvalidReorder)
	}

	if _, err := sc.supplierStorage.UpdatePositions(ctx, ids); err != nil {
		return utils.WrapIfErr(op, err)
	}

	return nil
}
