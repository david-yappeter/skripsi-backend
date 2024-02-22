package model

type TiktokPlatformAttribute struct {
	Id     string                 `db:"-"`
	Name   string                 `db:"-"`
	Values []TiktokAttributeValue `db:"-"`
}

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
