package model

const SequenceTableName = "sequences"

type Sequence struct {
	Id               string `db:"id"`
	UniqueIdentifier string `db:"unique_identifier"`
	Sequence         int    `db:"sequence"`
	Timestamp
}

func (m *Sequence) TableName() string {
	return SequenceTableName
}

func (m *Sequence) TableIds() []string {
	return []string{"id"}
}

func (m *Sequence) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":                m.Id,
		"unique_identifier": m.UniqueIdentifier,
		"sequence":          m.Sequence,
		"created_at":        m.CreatedAt,
		"updated_at":        m.UpdatedAt,
	}
}
