package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrOrderDeleteProtect = errors.New("service.Order: delete protection rows limit")
)

type OrderStorager interface {
	GetCounter(context.Context) (int64, error)
	FindMany(context.Context, storage.FindManyParams) ([]model.OrderHeader, int64, error)
	FindManyLinesByOrderID(context.Context, int64, storage.FindManyParams) ([]model.OrderLine, int64, model.OrderLineSummary, error)
	UpdateByID(context.Context, model.OrderHeader) error
	DeleteManyByID(context.Context, []int64) (int64, error)
}

type Order struct {
	orderStorage OrderStorager
}

func NewOrder(orderStorage OrderStorager) *Order {
	return &Order{orderStorage: orderStorage}
}

func (sc *Order) FetchCounter(ctx context.Context) (int64, error) {
	const op = "service.Order.FetchCounter"
	count, err := sc.orderStorage.GetCounter(ctx)
	return count, utils.WrapIfErr(op, err)
}

func (sc *Order) Fetch(ctx context.Context, params storage.FindManyParams) ([]model.OrderHeader, int64, error) {
	const op = "service.Order.Fetch"
	rows, count, err := sc.orderStorage.FindMany(ctx, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *Order) FetchLines(ctx context.Context, orderID int64, params storage.FindManyParams) ([]model.OrderLine, int64, model.OrderLineSummary, error) {
	const op = "service.Order.FetchLines"
	rows, count, summary, err := sc.orderStorage.FindManyLinesByOrderID(ctx, orderID, params)
	return rows, count, summary, utils.WrapIfErr(op, err)
}

func (sc *Order) SaveForm(ctx context.Context, dto model.OrderHeader) error {
	const op = "service.Order.SaveForm"
	err := sc.orderStorage.UpdateByID(ctx, dto)
	return utils.WrapIfErr(op, err)
}

func (sc *Order) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.Order.DeleteListID"
	if len(ids) > 20 {
		return 0, utils.WrapIfErr(op, ErrOrderDeleteProtect)
	}
	affected, err := sc.orderStorage.DeleteManyByID(ctx, ids)
	return affected, utils.WrapIfErr(op, err)
}
