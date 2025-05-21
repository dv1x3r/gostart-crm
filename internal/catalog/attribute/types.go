package attribute

import "github.com/dv1x3r/w2go/w2"

type AttributeGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AttributeSet struct {
	ID               int                 `json:"id"`
	AttributeGroupID int                 `json:"attribute_group_id"`
	Name             w2.Editable[string] `json:"name"`
}

func (as *AttributeSet) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&as.ID,
		&as.AttributeGroupID,
		&as.Name,
	)
}

type AttributeValue struct {
	ID                int                 `json:"id"`
	AttributeSetID    int                 `json:"attribute_set_id"`
	Name              w2.Editable[string] `json:"name"`
	RelatedProducts   w2.Editable[int]    `json:"related_products"`
	PublishedProducts w2.Editable[int]    `json:"published_products"`
}

func (av *AttributeValue) ScanRow(scan func(dest ...any) error) error {
	return scan(
		&av.ID,
		&av.AttributeSetID,
		&av.Name,
		&av.RelatedProducts,
		&av.PublishedProducts,
	)
}
