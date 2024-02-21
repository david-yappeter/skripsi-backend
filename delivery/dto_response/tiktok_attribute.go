package dto_response

import "myapp/model"

type TiktokAttributeResponse struct {
	Id                  string                         `json:"id"`
	Name                string                         `json:"name"`
	IsCustomizable      bool                           `json:"is_customizable"`
	IsMultipleSelection bool                           `json:"is_multiple_selection"`
	Values              []TiktokAttributeValueResponse `json:"values"`
} // @name TiktokAttributeResponse

func NewTiktokAttributeResponse(tiktokAttribute model.TiktokAttribute) TiktokAttributeResponse {
	r := TiktokAttributeResponse{
		Id:                  tiktokAttribute.Id,
		Name:                tiktokAttribute.Name,
		IsCustomizable:      tiktokAttribute.IsCustomizable,
		IsMultipleSelection: tiktokAttribute.IsMultipleSelection,
		Values:              []TiktokAttributeValueResponse{},
	}

	for _, value := range tiktokAttribute.Values {
		r.Values = append(r.Values, NewTiktokAttributeValueResponse(value))
	}

	return r
}

type TiktokAttributeValueResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
} // @name TiktokAttributeValueResponse

func NewTiktokAttributeValueResponse(tiktokAttributeValue model.TiktokAttributeValue) TiktokAttributeValueResponse {
	r := TiktokAttributeValueResponse{
		Id:   tiktokAttributeValue.Id,
		Name: tiktokAttributeValue.Name,
	}

	return r
}
