package supplier

import (
	"context"

	"gostart-crm/internal/lib/db"

	"github.com/dv1x3r/w2go/w2"
)

type Repository interface {
	Store() db.Store
	FindMany(context.Context, w2.GridDataRequest) ([]Supplier, int, error)
	GetDropdown(context.Context, w2.DropdownRequest) ([]w2.DropdownValue, error)
	UpsertMany(context.Context, []Supplier) error
	DeleteManyByID(context.Context, []int) error
	UpdatePositions(context.Context, []int) error
}
