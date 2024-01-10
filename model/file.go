package model

import "myapp/data_type"

const FileTableName = "files"

type File struct {
	Id   string             `db:"id"`
	Name string             `db:"name"`
	Type data_type.FileType `db:"type"`
	Path string             `db:"path"`

	Timestamp
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
