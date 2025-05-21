package attribute

import (
	"context"
	"database/sql"
	"errors"

	"gostart-crm/internal/lib/wrap"

	"github.com/dv1x3r/w2go/w2"
	"github.com/dv1x3r/w2go/w2sort"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) FetchAllGroups(ctx context.Context) ([]AttributeGroup, error) {
	const op = "attribute.Service.FetchAllGroups"
	rows, err := s.repo.FindAllGroups(ctx)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) FetchGroupDropdown(ctx context.Context, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "attribute.Service.FetchGroupDropdown"
	rows, err := s.repo.GetGroupDropdown(ctx, r)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) DeleteGroup(ctx context.Context, id int) error {
	const op = "attribute.Service.DeleteGroup"
	err := s.repo.DeleteGroupByID(ctx, id)
	if s.repo.Store().IsErrConstraintTrigger(err) {
		return wrap.IfErr(op, ErrAttributeHasRelated)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) SaveGroup(ctx context.Context, dto AttributeGroup) error {
	const op = "attribute.Service.SaveGroup"
	err := s.repo.UpsertGroup(ctx, dto)
	if s.repo.Store().IsErrConstraintUnique(err) {
		return wrap.IfErr(op, ErrAttributeExists)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) FetchSets(ctx context.Context, groupID int, r w2.GridDataRequest) ([]AttributeSet, int, error) {
	const op = "attribute.Service.FetchSets"
	rows, count, err := s.repo.FindManySetsByGroupID(ctx, groupID, r)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) SaveSetChanges(ctx context.Context, groupID int, changes []AttributeSet) error {
	const op = "attribute.Service.SaveSetChanges"
	err := s.repo.UpsertManySets(ctx, groupID, changes)
	if s.repo.Store().IsErrConstraintUnique(err) {
		return wrap.IfErr(op, ErrAttributeExists)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) DeleteSetListID(ctx context.Context, ids []int) error {
	const op = "attribute.Service.DeleteSetListID"
	err := s.repo.DeleteManySetsByID(ctx, ids)
	if s.repo.Store().IsErrConstraintTrigger(err) {
		return wrap.IfErr(op, ErrAttributeHasRelated)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) ReorderSets(ctx context.Context, r w2.GridReorderRequest) error {
	const op = "attribute.Service.ReorderSets"

	groupID, err := s.repo.GetGroupIDBySetID(ctx, r.RecID)
	if errors.Is(err, sql.ErrNoRows) {
		return wrap.IfErr(op, ErrAttributeNotFound)
	} else if err != nil {
		return wrap.IfErr(op, err)
	}

	sets, _, err := s.repo.FindManySetsByGroupID(ctx, groupID, w2.GridDataRequest{})
	if err != nil {
		return wrap.IfErr(op, err)
	}

	ids := make([]int, len(sets))
	for i, set := range sets {
		ids[i] = set.ID
	}

	if err := w2sort.ReorderArray(ids, r); err != nil {
		return wrap.IfErr(op, ErrAttributeInvalidReorder)
	}

	return wrap.IfErr(op, s.repo.UpdateSetPositions(ctx, ids))
}

func (s *Service) FetchValues(ctx context.Context, setID int, r w2.GridDataRequest) ([]AttributeValue, int, error) {
	const op = "attribute.Service.FetchValues"
	rows, count, err := s.repo.FindManyValuesBySetID(ctx, setID, r)
	return rows, count, wrap.IfErr(op, err)
}

func (s *Service) FetchValueDropdownBySetID(ctx context.Context, setID int, r w2.DropdownRequest) ([]w2.DropdownValue, error) {
	const op = "attribute.Service.FetchValueDropdownBySetID"
	rows, err := s.repo.GetValueDropdownBySetID(ctx, setID, r)
	return rows, wrap.IfErr(op, err)
}

func (s *Service) SaveValueChanges(ctx context.Context, setID int, changes []AttributeValue) error {
	const op = "attribute.Service.SaveValueChanges"
	err := s.repo.UpsertManyValues(ctx, setID, changes)
	if s.repo.Store().IsErrConstraintUnique(err) {
		return wrap.IfErr(op, ErrAttributeExists)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) DeleteValueListID(ctx context.Context, ids []int) error {
	const op = "attribute.Service.DeleteValueListID"
	err := s.repo.DeleteManyValuesByID(ctx, ids)
	if s.repo.Store().IsErrConstraintTrigger(err) {
		return wrap.IfErr(op, ErrAttributeHasRelated)
	}
	return wrap.IfErr(op, err)
}

func (s *Service) ReorderValues(ctx context.Context, r w2.GridReorderRequest) error {
	const op = "attribute.Service.ReorderValues"

	setID, err := s.repo.GetSetIDByValueID(ctx, r.RecID)
	if errors.Is(err, sql.ErrNoRows) {
		return wrap.IfErr(op, ErrAttributeNotFound)
	} else if err != nil {
		return wrap.IfErr(op, err)
	}

	values, _, err := s.repo.FindManyValuesBySetID(ctx, setID, w2.GridDataRequest{})
	if err != nil {
		return wrap.IfErr(op, err)
	}

	ids := make([]int, len(values))
	for i, value := range values {
		ids[i] = value.ID
	}

	if err := w2sort.ReorderArray(ids, r); err != nil {
		return wrap.IfErr(op, ErrAttributeInvalidReorder)
	}

	return wrap.IfErr(op, s.repo.UpdateValuePositions(ctx, ids))
}
