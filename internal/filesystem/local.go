package filesystem

import (
	"context"
	"fmt"
	"io"
	"myapp/util"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	LocalMainPath = "public"
	LocalTmpPath  = "tmp"
)

type localClient struct {
	config    *LocalClientConfig
	chunkSize int
}

type LocalClientConfig struct {
	BasePath string
	BaseUrl  string
}

func (l *localClient) fullpath(path string) (string, error) {
	path = util.ParsePathTrim(path)

	if strings.HasPrefix(path, "./") {
		return "", ErrPathIsNotAllowed
	}

	if regexp.MustCompile(`(\.\.\/)+`).MatchString(path) {
		return "", ErrPathIsNotAllowed
	}

	if strings.HasSuffix(path, "/..") {
		return "", ErrPathIsNotAllowed
	}

	return fmt.Sprintf("%s/%s", l.config.BasePath, path), nil
}

func (l *localClient) Delete(path string) error {
	fullpath, err := l.fullpath(path)
	if err != nil {
		return err
	}

	return os.Remove(fullpath)
}

func (l *localClient) DeleteFolderContents(folderPath string) error {
	folderFullpath, err := l.fullpath(folderPath)
	if err != nil {
		return err
	}

	fi, err := os.Stat(folderFullpath)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return ErrPathIsNotDirectory
	}

	return os.RemoveAll(folderFullpath)
}

func (l *localClient) Has(path string) (bool, error) {
	fullpath, err := l.fullpath(path)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(fullpath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (l *localClient) Open(path string) (io.ReadSeekCloser, error) {
	fullpath, err := l.fullpath(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotExist
		}

		return nil, err
	}

	return f, err
}

func (l *localClient) ReadFile(path string) ([]byte, error) {
	fullpath, err := l.fullpath(path)
	if err != nil {
		return nil, err
	}

	f, err := os.ReadFile(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotExist
		}

		return nil, err
	}

	return f, err
}

func (l *localClient) Stream(ctx context.Context, path string) (ReadCloserWithContent, error) {
	fullpath, err := l.fullpath(path)
	if err != nil {
		return nil, err
	}

	fStream, err := newFileStream(fullpath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotExist
		}

		return nil, err
	}

	return fStream, err
}

func (l *localClient) Url(path string) string {
	return fmt.Sprintf("%s/%s", l.config.BaseUrl, util.ParsePathTrim(path))
}

func (l *localClient) CopyTo(ctx context.Context, fromPath string, toPath string) error {
	file, err := l.Open(fromPath)
	if err != nil {
		return err
	}

	if err := l.Write(ctx, file, toPath); err != nil {
		return err
	}

	return nil
}

func (l *localClient) Write(ctx context.Context, r io.ReadSeekCloser, path string) error {
	var (
		chunk = make([]byte, l.chunkSize)
		n     int
		err   error
	)

	fullpath, err := l.fullpath(path)
	if err != nil {
		return err
	}

	// to make sure folder already exist
	os.MkdirAll(filepath.Dir(fullpath), os.ModePerm)

	f, err := os.Create(fullpath)
	if err != nil {
		return err
	}
	defer f.Close()

	for {
		select {
		case <-ctx.Done():
			err = ctx.Err()
			goto END

		default:
			n, err = r.Read(chunk)
			if err != nil && err != io.EOF {
				goto END
			}
			if n == 0 {
				err = nil
				goto END
			}

			if _, err = f.Write(chunk[:n]); err != nil {
				goto END
			}
		}
	}

END:
	if err != nil {
		l.Delete(path)
	}

	return err
}

func NewLocal(config *LocalClientConfig) Client {
	return &localClient{
		config:    config,
		chunkSize: defaultChunkSize,
	}
}
