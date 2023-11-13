package filesystem

import (
	"fmt"
	"myapp/global"
	"strings"
)

const (
	FilesystemLocal = "local"
	FilesystemAwsS3 = "aws_s3"
)

type FilesystemManager interface {
	Main() Client
	Tmp() Client
}

type filesystemManager struct {
	main Client
	tmp  Client
}

type Config struct {
	Filesystem string
}

func (m *filesystemManager) Main() Client {
	return m.main
}

func (m *filesystemManager) Tmp() Client {
	return m.tmp
}

func generateLocalClientConfig(localPath string) *LocalClientConfig {
	config := LocalClientConfig{
		BasePath: fmt.Sprintf("%s/%s", global.GetStorageDir(), localPath),
		BaseUrl:  fmt.Sprintf("%s/%s", strings.TrimRight(global.GetConfig().Uri, "/"), localPath),
	}
	return &config
}

func NewFilesystemManager(config Config) FilesystemManager {
	var main Client
	switch config.Filesystem {
	default:
		main = NewLocal(generateLocalClientConfig(LocalMainPath))
	}

	return &filesystemManager{
		main: main,
		tmp:  NewLocal(generateLocalClientConfig(LocalTmpPath)),
	}
}
