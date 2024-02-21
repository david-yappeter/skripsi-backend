package model

type TiktokAttribute struct {
	Id                  string                 `db:"-"`
	Name                string                 `db:"-"`
	IsCustomizable      bool                   `db:"-"`
	IsMultipleSelection bool                   `db:"-"`
	Values              []TiktokAttributeValue `db:"-"`
}

type TiktokAttributeValue struct {
	Id   string `db:"-"`
	Name string `db:"-"`
}
