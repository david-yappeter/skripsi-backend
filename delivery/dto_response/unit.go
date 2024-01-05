package dto_response

import "myapp/model"

type UnitResponse struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description" extensions:"x-nullable"`
	Timestamp
} // @name UnitResponse

func NewUnitResponse(unit model.Unit) UnitResponse {
	r := UnitResponse{
		Id:          unit.Id,
		Name:        unit.Name,
		Description: unit.Description,
		Timestamp:   Timestamp(unit.Timestamp),
	}

	return r
}
