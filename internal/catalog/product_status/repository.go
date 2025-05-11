package productstatus

import (
	"context"

	"gostart-crm/internal/lib/db"

	"github.com/dv1x3r/w2go/w2"
)

type Repository interface {
	Store() db.Store
	FindMany(context.Context, w2.GridDataRequest) ([]ProductStatus, int, error)
	FindAll(context.Context) ([]ProductStatus, error)
	GetDropdown(context.Context, w2.DropdownRequest) ([]ProductStatusDropdown, error)
	UpsertMany(context.Context, []ProductStatus) error
	DeleteManyByID(context.Context, []int) error
	UpdatePositions(context.Context, []int) error
}
