package manager

import (
	"myapp/global"
	"myapp/infrastructure"
	filesystemInternal "myapp/internal/filesystem"
	jwtInternal "myapp/internal/jwt"

	"myapp/repository"
	"myapp/use_case"
)

type Config int

const (
	LoadDefault Config = 1 << iota
	LoadUseCase
)

const DefaultConfig = LoadDefault | LoadUseCase

type Container struct {
	config Config

	// manager
	infrastructureManager infrastructure.InfrastructureManager
	repositoryManager     repository.RepositoryManager
	useCaseManager        use_case.UseCaseManager
	filesystemManager     filesystemInternal.FilesystemManager
	jwt                   jwtInternal.Jwt
}

func (c *Container) withUseCase() bool {
	return c.config&LoadUseCase != 0
}

func (c *Container) RepositoryManager() repository.RepositoryManager {
	return c.repositoryManager
}

func (c *Container) UseCaseManager() use_case.UseCaseManager {
	return c.useCaseManager
}

func (c *Container) LoggerStack() infrastructure.LoggerStack {
	return c.infrastructureManager.GetLoggerStack()
}

func (c *Container) MigrateDB(isRollingBack bool, steps int) error {
	return c.infrastructureManager.MigrateDB(isRollingBack, steps)
}

func (c *Container) RefreshDB() error {
	if err := c.infrastructureManager.RefreshDB(); err != nil {
		return err
	}

	if err := c.Close(); err != nil {
		return err
	}

	*c = *NewContainer(c.config)

	return nil
}

func (c Container) Close() error {
	if err := c.infrastructureManager.CloseDB(); err != nil {
		return err
	}

	return nil
}

func NewContainer(config Config) *Container {
	container := &Container{
		config: config,
	}

	container.infrastructureManager = infrastructure.NewInfrastructureManager()
	container.repositoryManager = repository.NewRepositoryManager(container.infrastructureManager)

	if container.withUseCase() {
		filesystemConfig := filesystemInternal.Config{
			Filesystem: global.GetFilesystem(),
		}
		container.filesystemManager = filesystemInternal.NewFilesystemManager(filesystemConfig)

		container.jwt = jwtInternal.NewJwt(
			global.GetJwtPrivateKeyFilePath(),
			global.GetJwtPublicKeyFilePath(),
		)

		container.useCaseManager = use_case.NewUseCaseManager(
			container.repositoryManager,
			container.filesystemManager,
			container.jwt,
			container.LoggerStack(),
		)
	}

	return container
}
