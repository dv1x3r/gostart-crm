package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrOrderStatusExists         = errors.New("service.OrderStatus: already exists")
	ErrOrderStatusHasRelated     = errors.New("service.OrderStatus: has related objects")
	ErrOrderStatusInvalidReorder = errors.New("service.OrderStatus: invalid reorder parameters")
)

type OrderStatusStorager interface {
	FindMany(context.Context, storage.FindManyParams) ([]model.OrderStatus, int64, error)
	GetDropdown(context.Context, string, int) ([]model.W2Dropdown, error)
	UpsertMany(context.Context, []model.OrderStatus) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	UpdatePositions(context.Context, []int64) (int64, error)
}

type OrderStatus struct {
	orderStatusStorage OrderStatusStorager
}

func NewOrderStatus(orderStatusStorage OrderStatusStorager) *OrderStatus {
	return &OrderStatus{orderStatusStorage: orderStatusStorage}
}

func (sc *OrderStatus) Fetch(ctx context.Context, params storage.FindManyParams) ([]model.OrderStatus, int64, error) {
	const op = "service.OrderStatus.Fetch"
	rows, count, err := sc.orderStatusStorage.FindMany(ctx, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *OrderStatus) FetchDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "service.OrderStatus.FetchDropdown"
	rows, err := sc.orderStatusStorage.GetDropdown(ctx, search, max)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *OrderStatus) SavePartials(ctx context.Context, dtos []model.OrderStatus) (int64, error) {
	const op = "service.OrderStatus.SavePartials"
	affected, err := sc.orderStatusStorage.UpsertMany(ctx, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrOrderStatusExists, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *OrderStatus) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.OrderStatus.DeleteListID"
	affected, err := sc.orderStatusStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrOrderStatusHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *OrderStatus) Reorder(ctx context.Context, objectID int64, moveBefore int64) error {
	const op = "service.OrderStatus.Reorder"

	group, _, err := sc.orderStatusStorage.FindMany(ctx, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	ids := make([]int64, len(group))
	for i, value := range group {
		ids[i] = value.ID
	}

	if ok := utils.ReorderBefore(ids, objectID, moveBefore); !ok {
		return utils.WrapIfErr(op, ErrOrderStatusInvalidReorder)
	}

	if _, err := sc.orderStatusStorage.UpdatePositions(ctx, ids); err != nil {
		return utils.WrapIfErr(op, err)
	}

	return nil
}
