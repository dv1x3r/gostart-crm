package category

import (
	"context"

	"gostart-crm/internal/lib/db"

	"github.com/dv1x3r/w2go/w2"
)

type Repository interface {
	Store() db.Store
	GetTree(context.Context, bool) (Category, error)
	GetDropdown(context.Context, w2.DropdownRequest, bool) ([]w2.DropdownValue, error)
	GetParentIDByCategoryID(context.Context, int) (int, error)
	FindManyByParentID(context.Context, int, w2.GridDataRequest) ([]Category, int, error)
	DeleteManyByID(context.Context, []int) error
	UpsertMany(context.Context, []Category) error
	UpdatePositions(context.Context, []int) error
}
