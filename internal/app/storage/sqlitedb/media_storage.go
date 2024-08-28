package sqlitedb

import (
	"context"

	"gostart-crm/internal/app/model"
	"gostart-crm/internal/app/utils"

	"github.com/jmoiron/sqlx"
)

type Media struct {
	db *sqlx.DB
}

func NewMedia(db *sqlx.DB) *Media {
	return &Media{db: db}
}

func (st *Media) GetProductIDs(ctx context.Context) ([]int, error) {
	const op = "sqlitedb.Media.GetProductIDs"
	query := "select distinct product_id from product_media"
	ids, err := runSelect[int](ctx, st.db, query, nil)
	return ids, utils.WrapIfErr(op, err)
}

func (st *Media) GetProductMediaFiles(ctx context.Context, productID int) ([]model.MediaFile, error) {
	const op = "sqlitedb.Media.GetProductMediaFiles"
	query := "select name, file, thumbnail from product_media where product_id = ?"
	rows, err := runSelect[model.MediaFile](ctx, st.db, query, []any{productID})
	return rows, utils.WrapIfErr(op, err)
}
