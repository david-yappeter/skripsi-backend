package model

type TiktokCategory struct {
	Id                 string            `db:"-"`
	ParentId           *string           `db:"-"`
	Name               string            `db:"-"`
	IsLeaf             bool              `db:"-"`
	ChildrenCategories []*TiktokCategory `db:"-"`
	ParentCategory     *TiktokCategory   `db:"-"`
}
