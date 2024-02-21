package dto_response

import "myapp/model"

type TiktokCategoryRuleResponse struct {
	PackageDimensionIsRequired bool `json:"package_dimension_is_required"`
	SizeChartIsSupported       bool `json:"size_chart_is_supported"`
	SizeChartIsRequired        bool `json:"size_chart_is_required"`
} // @name TiktokCategoryRuleResponse

func NewTiktokCategoryRuleResponse(tiktokCategoryRule model.TiktokCategoryRule) TiktokCategoryRuleResponse {
	r := TiktokCategoryRuleResponse{
		PackageDimensionIsRequired: tiktokCategoryRule.PackageDimensionIsRequired,
		SizeChartIsSupported:       tiktokCategoryRule.SizeChartIsSupported,
		SizeChartIsRequired:        tiktokCategoryRule.SizeChartIsRequired,
	}

	return r
}
