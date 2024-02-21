package dto_response

import "myapp/model"

type TiktokBrandResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
} // @name TiktokBrandResponse

func NewTiktokBrandResponse(tiktokBrand model.TiktokBrand) TiktokBrandResponse {
	r := TiktokBrandResponse{
		Id:   tiktokBrand.Id,
		Name: tiktokBrand.Name,
	}

	return r
}
