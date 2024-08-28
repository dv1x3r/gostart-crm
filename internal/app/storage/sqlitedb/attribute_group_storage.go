package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/storage"
	"gostart-crm/internal/app/utils"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
)

type AttributeGroup struct {
	db *sqlx.DB
}

func NewAttributeGroup(db *sqlx.DB) *AttributeGroup {
	return &AttributeGroup{db: db}
}

func (st *AttributeGroup) getQueryFindAll() (string, []any) {
	sb := sqlbuilder.Select("id", "name")
	sb.From("attribute_group")
	sb.OrderBy("name")
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeGroup) FindAll(ctx context.Context) ([]model.AttributeGroup, error) {
	const op = "sqlitedb.AttributeGroup.FindAll"
	query, args := st.getQueryFindAll()
	rows, err := runSelect[model.AttributeGroup](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *AttributeGroup) getQueryGetDropdown(search string, max int) (string, []any) {
	sb := sqlbuilder.Select("id", "name as text")
	sb.From("attribute_group")
	if search != "" {
		sb.Where(sb.Like("name", "%"+search+"%"))
	}
	sb.OrderBy("name")
	sb.Limit(max)
	return sb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeGroup) GetDropdown(ctx context.Context, search string, max int) ([]model.W2Dropdown, error) {
	const op = "sqlitedb.AttributeGroup.GetDropdown"
	query, args := st.getQueryGetDropdown(search, max)
	rows, err := runSelect[model.W2Dropdown](ctx, st.db, query, args)
	return rows, utils.WrapIfErr(op, err)
}

func (st *AttributeGroup) getQueryInsert(dto model.AttributeGroup) (string, []any) {
	ib := sqlbuilder.InsertInto("attribute_group")
	ib.Cols("name")
	ib.Values(dto.Name)
	return ib.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeGroup) getQueryUpdateByID(dto model.AttributeGroup) (string, []any) {
	ub := sqlbuilder.Update("attribute_group")
	ub.Where(ub.EQ("id", dto.ID))
	ub.SetMore("updated_at = unixepoch()")
	storage.ApplyUpdateSetPartial(ub, dto.Partial, "Name", "name", dto.Name)
	return ub.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeGroup) UpsertOne(ctx context.Context, dto model.AttributeGroup) error {
	const op = "sqlitedb.AttributeGroup.UpsertOne"
	if dto.ID == 0 {
		query, args := st.getQueryInsert(dto)
		_, err := runExecAffected(ctx, st.db, query, args)
		return utils.WrapIfErr(op, err)
	} else {
		query, args := st.getQueryUpdateByID(dto)
		_, err := runExecAffected(ctx, st.db, query, args)
		return utils.WrapIfErr(op, err)
	}
}

func (st *AttributeGroup) getQueryDeleteByID(id int64) (string, []any) {
	dlb := sqlbuilder.DeleteFrom("attribute_group")
	dlb.Where(dlb.EQ("id", id))
	return dlb.BuildWithFlavor(sqlbuilder.SQLite)
}

func (st *AttributeGroup) DeleteByID(ctx context.Context, id int64) error {
	const op = "sqlitedb.AttributeGroup.DeleteByID"
	query, args := st.getQueryDeleteByID(id)
	_, err := runExecAffected(ctx, st.db, query, args)
	return utils.WrapIfErr(op, err)
}
