package use_case

import (
	"context"
	"myapp/constant"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	validatorInternal "myapp/internal/gin/validator"
)

const (
	extensionTypeWord       = "word"
	extensionTypeExcel      = "excel"
	extensionTypePowerPoint = "powerpoint"
	extensionTypePdf        = "pdf"
	extensionTypeImage      = "image"
	extensionTypeGif        = "gif"
	extensionTypeAudio      = "audio"
	extensionTypeVideo      = "video"
	extensionTypeCompressed = "compressed"
	extensionTypeMedical    = "medical"
)

var (
	Validator validatorInternal.Validator = validatorInternal.New()

	extensions = map[string][]string{
		extensionTypeWord: {
			".docs",
			".doc",
			".docx",
		},

		extensionTypeExcel: {
			".xlsx",
			".xls",
			".xltx",
			".xlsb",
			".csv",
		},

		extensionTypePowerPoint: {
			".ppt",
			".pptx",
		},

		extensionTypePdf: {
			".pdf",
		},

		extensionTypeImage: {
			".jpeg",
			".jpg",
			".png",
			".jfif",
		},

		extensionTypeGif: {
			".gif",
		},

		extensionTypeAudio: {
			".mp3",
			".mpeg",
		},

		extensionTypeVideo: {
			".mp4",
		},

		extensionTypeCompressed: {
			".zip",
		},

		extensionTypeMedical: {
			".dcm",
			".dicom",
			".dicm",
			".DCM",
			".DICOM",
		},
	}
)

func listSupportedExtension(extensionTypes []string) []string {
	supportedExtensions := []string{}
	for _, extensionType := range extensionTypes {
		supportedExtensions = append(supportedExtensions, extensions[extensionType]...)
	}

	return supportedExtensions
}

type FilesystemCopy struct {
	Filesystem filesystem.Client
	Path       string
}

func (u FilesystemCopy) CopyTo(ctx context.Context, dest FilesystemCopy) error {
	if u.Filesystem == nil || u.Path == "" {
		panic("source filesystem and path must not empty")
	}

	if dest.Filesystem == nil || dest.Path == "" {
		panic("destination filesystem and path must not empty")
	}

	reader, err := u.Filesystem.Open(u.Path)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := dest.Filesystem.Write(ctx, reader, dest.Path); err != nil {
		return err
	}

	return nil
}

func (u FilesystemCopy) MustCopyTo(ctx context.Context, dest FilesystemCopy) {
	err := u.CopyTo(ctx, dest)
	if err != nil {
		panic(err)
	}
}

func panicIfErr(err error, excludedErrs ...error) {
	if err != nil {
		for _, excludedErr := range excludedErrs {
			if err == excludedErr {
				return
			}
		}
		panic(err)
	}
}

func panicIfRepositoryError(err error, errNoDataValidateMessage string, isValidate bool) {
	if err != nil {
		if err == constant.ErrNoData {
			if isValidate {
				panic(dto_response.NewBadRequestErrorResponse(errNoDataValidateMessage))
			}

			panic(dto_response.NewNotFoundErrorResponse("Data not found"))
		}

		panic(err)
	}
}
