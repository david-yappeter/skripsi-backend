package dto_response

import "myapp/model"

type TiktokPlatformImageResponse struct {
	Height   int    `json:"height"`
	Width    int    `json:"width"`
	ThumbUrl string `json:"thumb_url"`
	Uri      string `json:"uri"`
	Url      string `json:"url"`
} // @name TiktokPlatformImageResponse

func NewTiktokPlatformImageResponse(tiktokPlatformImage model.TiktokPlatformImage) TiktokPlatformImageResponse {
	r := TiktokPlatformImageResponse{
		Height:   tiktokPlatformImage.Height,
		Width:    tiktokPlatformImage.Width,
		ThumbUrl: tiktokPlatformImage.ThumbUrl,
		Uri:      tiktokPlatformImage.Uri,
		Url:      tiktokPlatformImage.Url,
	}

	return r
}
