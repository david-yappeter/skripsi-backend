package main

import (
	"context"
	"fmt"
	"myapp/cmd"
	"myapp/global"
	filesystemInternal "myapp/internal/filesystem"
	"myapp/manager"
	"myapp/use_case"
)

// @title		Mortal Health - Clinic Pilot API
// @version	0.0.1
// @host		cp-api.mortalhealth.com
// @BasePath	/
func main() {

	cmd.Execute()

	return
	filesystemConfig := filesystemInternal.Config{
		Filesystem: global.GetFilesystem(),
	}
	filesystemManager := filesystemInternal.NewFilesystemManager(filesystemConfig)

	container := manager.NewContainer(manager.DefaultConfig)
	defer func() {
		if err := container.Close(); err != nil {
			panic(err)
		}
	}()

	f, err := use_case.GenerateStockReport(
		context.Background(),
		container.RepositoryManager(),
		filesystemManager.Main(),
	)

	fmt.Printf("FILEEE : %+v\n\nerr: %+v\n\n", f, err)

}
