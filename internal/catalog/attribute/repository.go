package attribute

import (
	"context"

	"gostart-crm/internal/lib/db"

	"github.com/dv1x3r/w2go/w2"
)

type Repository interface {
	Store() db.Store

	// groups
	FindAllGroups(context.Context) ([]AttributeGroup, error)
	GetGroupDropdown(context.Context, w2.DropdownRequest) ([]w2.DropdownValue, error)
	UpsertGroup(context.Context, AttributeGroup) error
	DeleteGroupByID(context.Context, int) error

	// sets
	FindManySetsByGroupID(context.Context, int, w2.GridDataRequest) ([]AttributeSet, int, error)
	GetGroupIDBySetID(context.Context, int) (int, error)
	UpsertManySets(context.Context, int, []AttributeSet) error
	DeleteManySetsByID(context.Context, []int) error
	UpdateSetPositions(context.Context, []int) error

	// values
	FindManyValuesBySetID(context.Context, int, w2.GridDataRequest) ([]AttributeValue, int, error)
	GetValueDropdownBySetID(context.Context, int, w2.DropdownRequest) ([]w2.DropdownValue, error)
	GetSetIDByValueID(context.Context, int) (int, error)
	UpsertManyValues(context.Context, int, []AttributeValue) error
	DeleteManyValuesByID(context.Context, []int) error
	UpdateValuePositions(context.Context, []int) error
}
