package service

import (
	"context"
	"errors"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"
)

var (
	ErrAttributeGroupExists     = errors.New("service.AttributeGroup: already exists")
	ErrAttributeGroupNotFound   = errors.New("service.AttributeGroup: not found")
	ErrAttributeGroupHasRelated = errors.New("service.AttributeGroup: has related objects")
)

type AttributeGroupStorager interface {
	FindAll(context.Context) ([]model.AttributeGroup, error)
	GetDropdown(context.Context, string, int) ([]model.W2Dropdown, error)
	UpsertOne(context.Context, model.AttributeGroup) error
	DeleteByID(context.Context, int64) error
}

type AttributeGroup struct {
	attributeGroupStorage AttributeGroupStorager
}

func NewAttributeGroup(attributeGroupStorage AttributeGroupStorager) *AttributeGroup {
	return &AttributeGroup{attributeGroupStorage: attributeGroupStorage}
}

func (sc *AttributeGroup) FetchAll(ctx context.Context) ([]model.AttributeGroup, error) {
	const op = "service.AttributeGroup.FetchAll"
	rows, err := sc.attributeGroupStorage.FindAll(ctx)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *AttributeGroup) FetchDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "service.AttributeGroup.FetchDropdown"
	rows, err := sc.attributeGroupStorage.GetDropdown(ctx, search, max)
	return rows, utils.WrapIfErr(op, err)
}

func (sc *AttributeGroup) Delete(ctx context.Context, id int64) error {
	const op = "service.AttributeGroup.Delete"
	err := sc.attributeGroupStorage.DeleteByID(ctx, id)
	if errors.Is(err, storage.ErrConstraintTrigger) {
		return utils.WrapIfErr(op, errors.Join(ErrAttributeGroupHasRelated, err))
	}
	return utils.WrapIfErr(op, err)
}

func (sc *AttributeGroup) Save(ctx context.Context, dto model.AttributeGroup) error {
	const op = "service.Attribute.Save"
	err := sc.attributeGroupStorage.UpsertOne(ctx, dto)
	if errors.Is(err, storage.ErrConstraintUnique) {
		return utils.WrapIfErr(op, errors.Join(ErrAttributeGroupExists, err))
	}
	return utils.WrapIfErr(op, err)
}
