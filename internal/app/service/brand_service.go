package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrBrandExists     = errors.New("service.Brand: already exists")
	ErrBrandHasRelated = errors.New("service.Brand: has related objects")
)

type BrandStorager interface {
	FindMany(context.Context, storage.FindManyParams) ([]model.Brand, int64, error)
	GetDropdown(context.Context, string, int) ([]model.W2Dropdown, error)
	UpsertMany(context.Context, []model.Brand) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
}

type Brand struct {
	brandStorage BrandStorager
}

func NewBrand(brandStorage BrandStorager) *Brand {
	return &Brand{brandStorage: brandStorage}
}

func (sc *Brand) Fetch(ctx context.Context, params storage.FindManyParams) ([]model.Brand, int64, error) {
	const op = "service.Brand.Fetch"
	rows, count, err := sc.brandStorage.FindMany(ctx, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *Brand) FetchDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "service.Brand.FetchDropdown"
	rows, err := sc.brandStorage.GetDropdown(ctx, search, max)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *Brand) SavePartials(ctx context.Context, dtos []model.Brand) (int64, error) {
	const op = "service.Brand.SavePartials"
	affected, err := sc.brandStorage.UpsertMany(ctx, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrBrandExists, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *Brand) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.Brand.DeleteListID"
	affected, err := sc.brandStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrBrandHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}
