package model

type ProductMedia struct {
	ID           int64  `json:"id" db:"id"`
	ProductID    int64  `json:"product_id" db:"product_id"`
	Name         string `json:"name" db:"name"`
	HasThumbnail bool   `json:"-" db:"has_thumbnail"`

	File         []byte  `json:"-" db:"-"`
	FileURL      string  `json:"file_url" db:"-"`
	Thumbnail    []byte  `json:"-" db:"-"`
	ThumbnailURL *string `json:"thumbnail_url" db:"-"`
}

type ProductMediaW2GridResponse = W2GridDataResponse[ProductMedia, any]
