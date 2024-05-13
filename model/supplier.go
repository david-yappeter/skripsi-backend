package model

const SupplierTableName = "suppliers"

type Supplier struct {
	Id             string  `db:"id"`
	SupplierTypeId string  `db:"supplier_type_id"`
	Code           string  `db:"code"`
	Name           string  `db:"name"`
	IsActive       bool    `db:"is_active"`
	Address        string  `db:"address"`
	Phone          string  `db:"phone"`
	Email          *string `db:"email"`
	Description    *string `db:"description"`

	Timestamp

	SupplierType *SupplierType `db:"-"`
}

func (m *Supplier) TableName() string {
	return SupplierTableName
}

func (m *Supplier) TableIds() []string {
	return []string{"id"}
}

func (m *Supplier) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"id":               m.Id,
		"supplier_type_id": m.SupplierTypeId,
		"code":             m.Code,
		"name":             m.Name,
		"is_active":        m.IsActive,
		"address":          m.Address,
		"phone":            m.Phone,
		"email":            m.Email,
		"description":      m.Description,
		"created_at":       m.CreatedAt,
		"updated_at":       m.UpdatedAt,
	}
}

type SupplierQueryOption struct {
	QueryOption

	IsActive        *bool
	SupplierTypeIds []string
	Phrase          *string
}

var _ PrepareOption = &SupplierQueryOption{}

func (o *SupplierQueryOption) SetDefaultFields() {
	if len(o.Fields) == 0 {
		o.Fields = []string{"s.*"}
	}
}

func (o *SupplierQueryOption) SetDefaultSorts() {
	if len(o.Sorts) == 0 {
		o.Sorts = Sorts{
			{Field: "updated_at", Direction: "desc"},
		}
	}
}

type SupplierDebtSummary struct {
	SupplierId   string  `db:"supplier_id"`
	SupplierName string  `db:"supplier_name"`
	TotalDebt    float64 `db:"total_debt"`
}
