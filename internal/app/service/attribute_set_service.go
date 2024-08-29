package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrAttributeSetExists         = errors.New("service.AttributeSet: already exists")
	ErrAttributeSetNotFound       = errors.New("service.AttributeSet: not found")
	ErrAttributeSetHasRelated     = errors.New("service.AttributeSet: has related objects")
	ErrAttributeSetInvalidReorder = errors.New("service.AttributeSet: invalid reorder parameters")
)

type AttributeSetStorager interface {
	FindManyByGroupID(context.Context, int64, storage.FindManyParams) ([]model.AttributeSet, int64, error)
	GetByID(context.Context, int64) (model.AttributeSet, bool, error)
	UpsertMany(context.Context, int64, []model.AttributeSet) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	UpdatePositions(context.Context, []int64) (int64, error)
}

type AttributeSet struct {
	attributeSetStorage AttributeSetStorager
}

func NewAttributeSet(attributeSetStorage AttributeSetStorager) *AttributeSet {
	return &AttributeSet{attributeSetStorage: attributeSetStorage}
}

func (sc *AttributeSet) Fetch(ctx context.Context, groupID int64, params storage.FindManyParams) ([]model.AttributeSet, int64, error) {
	const op = "service.AttributeSet.Fetch"
	rows, count, err := sc.attributeSetStorage.FindManyByGroupID(ctx, groupID, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *AttributeSet) SavePartials(ctx context.Context, groupID int64, dtos []model.AttributeSet) (int64, error) {
	const op = "service.AttributeSet.SavePartials"
	affected, err := sc.attributeSetStorage.UpsertMany(ctx, groupID, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrAttributeSetExists, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *AttributeSet) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.AttributeSet.DeleteListID"
	affected, err := sc.attributeSetStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrAttributeSetHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *AttributeSet) Reorder(ctx context.Context, objectID int64, moveBefore int64) error {
	const op = "service.AttributeSet.Reorder"

	object, ok, err := sc.attributeSetStorage.GetByID(ctx, objectID)
	if err != nil {
		return utils.WrapIfErr(op, err)
	} else if !ok {
		return utils.WrapIfErr(op, ErrAttributeSetNotFound)
	}

	group, _, err := sc.attributeSetStorage.FindManyByGroupID(ctx, object.AttributeGroupID, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	ids := make([]int64, len(group))
	for i, set := range group {
		ids[i] = set.ID
	}

	if ok := utils.ReorderBefore(ids, objectID, moveBefore); !ok {
		return utils.WrapIfErr(op, ErrAttributeSetInvalidReorder)
	}

	if _, err := sc.attributeSetStorage.UpdatePositions(ctx, ids); err != nil {
		return utils.WrapIfErr(op, err)
	}

	return nil
}
