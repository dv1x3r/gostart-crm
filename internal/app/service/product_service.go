package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrProductExists            = errors.New("service.Product: already exists")
	ErrProductNotFound          = errors.New("service.Product: not found")
	ErrProductDeleteProtect     = errors.New("service.Product: delete protection rows limit")
	ErrProductAttributeNotFound = errors.New("service.Product: attribute not found")
)

type ProductStorager interface {
	FindMany(context.Context, storage.FindManyParams, int64) ([]model.Product, int64, error)
	GetByID(context.Context, int64) (model.Product, bool, error)
	UpsertOne(context.Context, model.Product) (int64, error)
	UpdateMany(context.Context, []model.Product) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	FindManyAttributesByProductID(context.Context, int64, storage.FindManyParams) ([]model.ProductAttribute, int64, error)
	UpsertManyAttributes(context.Context, int64, []model.ProductAttribute) (int64, error)
}

type Product struct {
	productStorage ProductStorager
}

func NewProduct(productStorage ProductStorager) *Product {
	return &Product{productStorage: productStorage}
}

func (sc *Product) Fetch(ctx context.Context, params storage.FindManyParams, categoryID int64) ([]model.Product, int64, error) {
	const op = "service.Product.Fetch"
	rows, count, err := sc.productStorage.FindMany(ctx, params, categoryID)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *Product) FetchByID(ctx context.Context, id int64) (model.Product, error) {
	const op = "service.Product.FetchByID"
	product, ok, err := sc.productStorage.GetByID(ctx, id)
	if err != nil {
		return product, utils.WrapIfErr(op, err)
	} else if !ok {
		return product, utils.WrapIfErr(op, ErrProductNotFound)
	}
	return product, nil
}

func (sc *Product) SaveForm(ctx context.Context, dto model.Product) (int64, error) {
	const op = "service.Product.SaveForm"

	id, err := sc.productStorage.UpsertOne(ctx, dto)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrProductExists, err))
	}

	return id, utils.WrapIfErr(op, err)
}

func (sc *Product) SavePartials(ctx context.Context, dtos []model.Product) (int64, error) {
	const op = "service.Product.SavePartials"
	affected, err := sc.productStorage.UpdateMany(ctx, dtos)
	return affected, utils.WrapIfErr(op, err)
}

func (sc *Product) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.Product.DeleteListID"
	if len(ids) > 50 {
		return 0, utils.WrapIfErr(op, ErrProductDeleteProtect)
	}
	affected, err := sc.productStorage.DeleteManyByID(ctx, ids)
	return affected, utils.WrapIfErr(op, err)
}

func (sc *Product) FetchAttributes(ctx context.Context, productID int64, params storage.FindManyParams) ([]model.ProductAttribute, int64, error) {
	const op = "service.Product.FetchAttributes"
	rows, count, err := sc.productStorage.FindManyAttributesByProductID(ctx, productID, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *Product) SaveAttributePartials(ctx context.Context, productID int64, dtos []model.ProductAttribute) (int64, error) {
	const op = "service.Product.SaveAttributePartials"
	affected, err := sc.productStorage.UpsertManyAttributes(ctx, productID, dtos)
	if errors.Is(err, storage.ErrConstraintForeignKey) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrProductAttributeNotFound, err))
	}
	return affected, utils.WrapIfErr(op, err)
}
