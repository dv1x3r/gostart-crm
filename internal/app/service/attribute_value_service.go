package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrAttributeValueExists         = errors.New("service.AttributeValue: already exists")
	ErrAttributeValueNotFound       = errors.New("service.AttributeValue: not found")
	ErrAttributeValueHasRelated     = errors.New("service.AttributeValue: has related objects")
	ErrAttributeValueInvalidReorder = errors.New("service.AttributeValue: invalid reorder parameters")
)

type AttributeValueStorager interface {
	FindManyBySetID(context.Context, int64, storage.FindManyParams) ([]model.AttributeValue, int64, error)
	GetDropdownBySetID(context.Context, int64, string, int) ([]model.W2Dropdown, error)
	GetByID(context.Context, int64) (model.AttributeValue, bool, error)
	UpsertMany(context.Context, int64, []model.AttributeValue) (int64, error)
	DeleteManyByID(context.Context, []int64) (int64, error)
	UpdatePositions(context.Context, []int64) (int64, error)
}

type AttributeValue struct {
	attributeValueStorage AttributeValueStorager
}

func NewAttributeValue(attributeValueStorage AttributeValueStorager) *AttributeValue {
	return &AttributeValue{attributeValueStorage: attributeValueStorage}
}

func (sc *AttributeValue) Fetch(ctx context.Context, setID int64, params storage.FindManyParams) ([]model.AttributeValue, int64, error) {
	const op = "service.AttributeValue.Fetch"
	rows, count, err := sc.attributeValueStorage.FindManyBySetID(ctx, setID, params)
	return rows, count, utils.WrapIfErr(op, err)
}

func (sc *AttributeValue) FetchDropdownBySetID(ctx context.Context, setID int64, search string, max int) ([]model.W2Dropdown, error) {
	const op = "service.AttributeValue.FetchDropdownBySetID"
	rows, err := sc.attributeValueStorage.GetDropdownBySetID(ctx, setID, search, max)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *AttributeValue) SavePartials(ctx context.Context, setID int64, dtos []model.AttributeValue) (int64, error) {
	const op = "service.AttributeValue.SavePartials"
	affected, err := sc.attributeValueStorage.UpsertMany(ctx, setID, dtos)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrAttributeValueExists, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *AttributeValue) DeleteListID(ctx context.Context, ids []int64) (int64, error) {
	const op = "service.Attribute.DeleteListID"
	affected, err := sc.attributeValueStorage.DeleteManyByID(ctx, ids)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return 0, utils.WrapIfErr(op, errors.Join(ErrAttributeValueHasRelated, err))
	}
	return affected, utils.WrapIfErr(op, err)
}

func (sc *AttributeValue) Reorder(ctx context.Context, objectID int64, moveBefore int64) error {
	const op = "service.AttributeValue.Reorder"

	object, ok, err := sc.attributeValueStorage.GetByID(ctx, objectID)
	if err != nil {
		return utils.WrapIfErr(op, err)
	} else if !ok {
		return utils.WrapIfErr(op, ErrAttributeValueNotFound)
	}

	group, _, err := sc.attributeValueStorage.FindManyBySetID(ctx, object.AttributeSetID, storage.FindManyParams{})
	if err != nil {
		return utils.WrapIfErr(op, err)
	}

	ids := make([]int64, len(group))
	for i, value := range group {
		ids[i] = value.ID
	}

	if ok := utils.ReorderBefore(ids, objectID, moveBefore); !ok {
		return utils.WrapIfErr(op, ErrAttributeValueInvalidReorder)
	}

	if _, err := sc.attributeValueStorage.UpdatePositions(ctx, ids); err != nil {
		return utils.WrapIfErr(op, err)
	}

	return nil
}
