package dto_response

import "myapp/model"

type TiktokCategoryResponse struct {
	Id                 string                   `json:"id"`
	Name               string                   `json:"name"`
	IsLeaf             bool                     `json:"is_leaf"`
	ChildrenCategories []TiktokCategoryResponse `json:"children_categories"`
} // @name TiktokCategoryResponse

func NewTiktokCategoryResponse(tiktokCategory model.TiktokCategory) TiktokCategoryResponse {
	r := TiktokCategoryResponse{
		Id:                 tiktokCategory.Id,
		Name:               tiktokCategory.Name,
		IsLeaf:             tiktokCategory.IsLeaf,
		ChildrenCategories: []TiktokCategoryResponse{},
	}

	if r.IsLeaf {
		r.ChildrenCategories = nil
	}

	for _, tiktokCategoryChildren := range tiktokCategory.ChildrenCategories {
		r.ChildrenCategories = append(r.ChildrenCategories, NewTiktokCategoryResponse(*tiktokCategoryChildren))
	}

	return r
}
