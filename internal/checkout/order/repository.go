package order

import (
	"context"

	"gostart-crm/internal/lib/db"

	"github.com/dv1x3r/w2go/w2"
)

type Repository interface {
	Store() db.Store
	FindMany(context.Context, w2.GridDataRequest) ([]Order, int, error)
	UpdateByID(context.Context, int, Order) error
	DeleteManyByID(context.Context, []int) error
}
