package product

import (
	"context"

	"gostart-crm/internal/lib/db"

	"github.com/dv1x3r/w2go/w2"
)

type Repository interface {
	Store() db.Store
	FindManyByCategoryID(context.Context, w2.GridDataRequest, int) ([]Product, int, error)
	GetByID(context.Context, int) (Product, error)
	UpsertOne(context.Context, Product) (int, error)
	UpdateMany(context.Context, []Product) error
	DeleteManyByID(context.Context, []int) error
	FindAvailableByCategoryID(context.Context, ListDataRequest, int) ([]Product, error)
	FindManyAttributesByProductID(context.Context, int, w2.GridDataRequest) ([]ProductAttribute, int, error)
	UpsertManyAttributes(context.Context, int, []ProductAttribute) error
}
