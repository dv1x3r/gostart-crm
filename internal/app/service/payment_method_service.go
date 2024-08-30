package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrPaymentMethodExists         = errors.New("service.PaymentMethod: already exists")
	ErrPaymentMethodHasRelated     = errors.New("service.PaymentMethod: has related objects")
	ErrPaymentMethodInvalidReorder = errors.New("service.PaymentMethod: invalid reorder parameters")
)

type PaymentMethodStorager interface {
	FindMany(context.Context, storage.FindManyParams) ([]model.PaymentMethod, int64, error)
	GetDropdown(context.Context, string, int) ([]model.W2Dropdown, error)
	UpsertMany(context.Context, []model.PaymentMethod) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	UpdatePositions(context.Context, []int64) (int64, error)
}

type PaymentMethod struct {
	paymentMethodStorage PaymentMethodStorager
}

func NewPaymentMethod(paymentMethodStorage PaymentMethodStorager) *PaymentMethod {
	return &PaymentMethod{paymentMethodStorage: paymentMethodStorage}
}

func (sc *PaymentMethod) Fetch(ctx context.Context, params storage.FindManyParams) ([]model.PaymentMethod, int64, error) {
	const op = "service.PaymentMethod.Fetch"
	rows, count, err := sc.paymentMethodStorage.FindMany(ctx, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *PaymentMethod) FetchDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "service.PaymentMethod.FetchDropdown"
	rows, err := sc.paymentMethodStorage.GetDropdown(ctx, search, max)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *PaymentMethod) SavePartials(ctx context.Context, dtos []model.PaymentMethod) (int64, error) {
	const op = "service.PaymentMethod.SavePartials"
	affected, err := sc.paymentMethodStorage.UpsertMany(ctx, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrPaymentMethodExists, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *PaymentMethod) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.PaymentMethod.DeleteListID"
	affected, err := sc.paymentMethodStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrPaymentMethodHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *PaymentMethod) Reorder(ctx context.Context, objectID int64, moveBefore int64) error {
	const op = "service.PaymentMethod.Reorder"

	group, _, err := sc.paymentMethodStorage.FindMany(ctx, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	ids := make([]int64, len(group))
	for i, value := range group {
		ids[i] = value.ID
	}

	if ok := utils.ReorderBefore(ids, objectID, moveBefore); !ok {
		return utils.WrapIfErr(op, ErrPaymentMethodInvalidReorder)
	}

	if _, err := sc.paymentMethodStorage.UpdatePositions(ctx, ids); err != nil {
		return utils.WrapIfErr(op, err)
	}

	return nil
}
