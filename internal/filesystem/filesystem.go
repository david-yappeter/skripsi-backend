package filesystem

import (
	"context"
	"io"
	"os"
)

const defaultChunkSize = 5 << 20 // 5 MB

type Client interface {
	Delete(path string) error
	DeleteFolderContents(folderPath string) error
	Has(path string) (bool, error)
	Open(path string) (io.ReadSeekCloser, error)
	ReadFile(path string) ([]byte, error)
	Stream(ctx context.Context, path string) (ReadCloserWithContent, error)
	Url(path string) string
	CopyTo(ctx context.Context, fromPath string, toPath string) error
	Write(ctx context.Context, r io.ReadSeekCloser, path string) error
}

type ReadCloserWithContent interface {
	ContentLength() int64
	ContentType() string
	io.ReadCloser
}

type ioStream struct {
	contentLength int64
	contentType   string
	pr            *io.PipeReader
	pw            *io.PipeWriter
	err           error
}

func (f *ioStream) closeWriter() error {
	return f.pw.Close()
}

func (f *ioStream) ContentLength() int64 {
	return f.contentLength
}

func (f *ioStream) ContentType() string {
	return f.contentType
}

func (f *ioStream) Read(p []byte) (n int, err error) {
	if f.err != nil {
		return 0, err
	}

	return f.pr.Read(p)
}

func (f *ioStream) Close() error {
	return f.pr.Close()
}

func (f *ioStream) WriteAt(p []byte, offset int64) (n int, err error) {
	// ignore 'offset' because we forced sequential downloads
	return f.pw.Write(p)
}

type fileStream struct {
	contentLength int64
	contentType   string
	file          *os.File
}

func (f *fileStream) ContentLength() int64 {
	return f.contentLength
}

func (f *fileStream) ContentType() string {
	return f.contentType
}

func (f *fileStream) Read(p []byte) (n int, err error) {
	return f.file.Read(p)
}

func (f *fileStream) Close() error {
	return f.file.Close()
}

func newIoStream(contentLength int64, contentType string) *ioStream {
	pr, pw := io.Pipe()

	return &ioStream{
		contentLength: contentLength,
		contentType:   contentType,
		pr:            pr,
		pw:            pw,
	}
}

func newFileStream(path string) (*fileStream, error) {
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotExist
		}

		return nil, err
	}

	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	contentType, err := guessMimeType(f)
	if err != nil {
		return nil, err
	}

	fStream := &fileStream{
		contentLength: fi.Size(),
		contentType:   contentType,
		file:          f,
	}

	return fStream, nil
}
