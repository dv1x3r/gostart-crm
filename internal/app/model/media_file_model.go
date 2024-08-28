package model

type MediaFile struct {
	Name      string `db:"name"`
	File      []byte `db:"file"`
	Thumbnail []byte `db:"thumbnail"`
}
