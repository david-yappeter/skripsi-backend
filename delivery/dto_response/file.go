package dto_response

import (
	"myapp/model"
)

type FileResponse struct {
	Id   string `json:"id" example:"f55824e5-7eec-43f9-90b5-cbddb7917c42" format:"uuid"`
	Name string `json:"name" example:"file_name.jpg"`
	Path string `json:"path" example:"/path/to/file_name.jpg"`
	Link string `json:"link" example:"https://example.com/path/to/file_name.jpg"`

	Timestamp
} // @name FileResponse

func NewFileResponse(file model.File) FileResponse {
	r := FileResponse{
		Id:        file.Id,
		Name:      file.Name,
		Path:      file.Path,
		Timestamp: Timestamp(file.Timestamp),
	}

	if file.Link != nil {
		r.Link = *file.Link
	}

	return r
}

func NewFileResponseP(file model.File) *FileResponse {
	r := NewFileResponse(file)
	return &r
}
