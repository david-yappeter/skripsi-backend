package use_case

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	"myapp/util"
	"path"
)

type baseFileUseCase struct {
	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client
}

func newBaseFileUseCase(
	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) baseFileUseCase {
	return baseFileUseCase{
		mainFilesystem: mainFilesystem,
		tmpFilesystem:  tmpFilesystem,
	}
}

type fileUploadTemporaryParams struct {
	supportedExtensions []string
	maxFileSizeInBytes  *int64
}

type fileUploadTemporaryToMainParams struct {
	deleteTmpOnSuccess bool
}

func (u *baseFileUseCase) mustValidateDocumentFileExtension(ext string, supportedExtensions []string) {
	if supportedExtensions == nil {
		return
	}

	if !util.StringInSlice(ext, supportedExtensions) {
		panic(dto_response.NewBadRequestErrorResponse("File extension is not supported"))
	}
}

func (u *baseFileUseCase) mustValidateFileSize(fileSize int64, maxFileSize *int64) {
	if fileSize == 0 {
		panic(dto_response.NewBadRequestErrorResponse("File size is 0B"))
	}

	if maxFileSize == nil || fileSize <= *maxFileSize {
		return
	}

	// human readable byte size format
	var (
		multiplier int64 = 1024
		humanize         = "Maximum file size is "
		divider    int64 = multiplier
		loop       int   = 0
	)

	// B
	if fileSize < multiplier {
		humanize += fmt.Sprintf("%d B", fileSize)
	} else {
		// KB MB GB TB PB EB
		for n := fileSize / multiplier; n >= multiplier; n /= multiplier {
			divider *= multiplier
			loop++
		}
		humanize += fmt.Sprintf("%.1f %cB", float64(fileSize)/float64(divider), "KMGTPE"[loop])
	}

	panic(dto_response.NewBadRequestErrorResponse(humanize))
}

func (u *baseFileUseCase) mustValidateTemporaryFilePathsExistance(filepaths []string) {
	for _, path := range filepaths {
		isExist, err := u.tmpFilesystem.Has(path)
		panicIfErr(err)
		if !isExist {
			if len(filepaths) > 1 {
				panic(dto_response.NewBadRequestErrorResponse("Some file not exist"))
			} else {
				panic(dto_response.NewBadRequestErrorResponse("File not exist"))
			}
		}
	}
}

func (u *baseFileUseCase) mustUploadFileToTemporary(ctx context.Context, entityName string, cloudFilename string, fileHeader *multipart.FileHeader, opt fileUploadTemporaryParams) string {
	ext := path.Ext(fileHeader.Filename)

	u.mustValidateDocumentFileExtension(ext, opt.supportedExtensions)
	u.mustValidateFileSize(fileHeader.Size, opt.maxFileSizeInBytes)

	path := fmt.Sprintf("%s/%s%s", entityName, util.NewKsuid(), cloudFilename)

	fileReader, err := fileHeader.Open()
	panicIfErr(err)
	defer fileReader.Close()

	panicIfErr(u.tmpFilesystem.Write(ctx, fileReader, path))

	return path
}

func (u *baseFileUseCase) mustUploadFileFromTemporaryToMain(ctx context.Context, entityName string, entityId string, cloudFilename string, tmpPath string, opt fileUploadTemporaryToMainParams) (string, string) {
	u.mustValidateTemporaryFilePathsExistance([]string{tmpPath})

	src := FilesystemCopy{
		Filesystem: u.tmpFilesystem,
		Path:       tmpPath,
	}

	dest := FilesystemCopy{
		Filesystem: u.mainFilesystem,
		Path:       fmt.Sprintf("%s/%s/%s", entityName, entityId, cloudFilename),
	}

	// TODO: how to implement errgroup.Go async to src.MustCopyTo()
	src.MustCopyTo(ctx, dest)

	if opt.deleteTmpOnSuccess {
		// TODO: how to implement delete later after use_case 'success' running, not in the middle of code
		go u.tmpFilesystem.Delete(tmpPath)
	}

	realCloudFileName := util.GetFilenameFromUploadPath(src.Path)

	return dest.Path, realCloudFileName
}

func (u *baseFileUseCase) mustUploadFileFromReaderToMain(ctx context.Context, entityName string, entityId string, cloudFilename string, reader io.ReadSeekCloser) string {
	destPath := fmt.Sprintf("%s/%s/%s", entityName, entityId, cloudFilename)

	if err := u.mainFilesystem.Write(ctx, reader, destPath); err != nil {
		panic(err)
	}

	return destPath
}

func (u baseFileUseCase) GetMainFilesystemLink(path string) string {
	return u.mainFilesystem.Url(path)
}
