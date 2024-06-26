package model

import (
	"myapp/data_type"
	"myapp/internal/filesystem"
	"myapp/util"
)

const FileTableName = "files"

type File struct {
	Id   string             `db:"id"`
	Name string             `db:"name"`
	Type data_type.FileType `db:"type"`
	Path string             `db:"path"`
	Timestamp

	Link *string `db:"-"`
}

func (m *File) TableName() string {
	return FileTableName
}

func (m *File) TableIds() []string {
	return []string{"id"}
}

func (m *File) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":         m.Id,
		"name":       m.Name,
		"type":       m.Type,
		"path":       m.Path,
		"created_at": m.CreatedAt,
		"updated_at": m.UpdatedAt,
	}
}

func (m *File) SetLink(fs filesystem.Client) {
	m.Link = util.StringP(fs.Url(m.Path))
}
