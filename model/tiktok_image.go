package model

type TiktokPlatformImage struct {
	Height   int    `db:"-"`
	Width    int    `db:"-"`
	ThumbUrl string `db:"-"`
	Uri      string `db:"-"`
	Url      string `db:"-"`
}
