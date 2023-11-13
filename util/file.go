package util

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
)

func GenerateFilepath(fileHeader *multipart.FileHeader) string {
	return fmt.Sprintf("%s_%s%s", fileHeader.Filename, NewKsuid(), filepath.Ext(fileHeader.Filename))
}

func GetFilenameFromUploadPath(path string) string {
	lengthOfKsuid := 27

	if strings.Contains(path, "/") {
		splitted := strings.Split(path, "/")
		filename := splitted[len(splitted)-1]

		return filename[lengthOfKsuid:]
	}
	return path[lengthOfKsuid:]
}
