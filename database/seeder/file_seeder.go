package seeder

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/model"
	"myapp/repository"
)

var (
	FileOne = model.File{
		Id:   "29e3a2e9-697d-4d74-8a35-8b8768e86cc2",
		Name: "Gambar Kaleng Cat Merah",
		Type: data_type.FileTypeProductUnitImage,
		Path: fmt.Sprintf("%s/2wag9wejGEqjgieay38EG8edg8e8s.jpg", constant.ProductUnitImagePath),
	}
)

func FileSeeder(repositoryManager repository.RepositoryManager) {
	fileRepository := repositoryManager.FileRepository()

	count, err := fileRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := fileRepository.InsertMany(context.Background(), getFileData()); err != nil {
		panic(err)
	}
}

func getFileData() []model.File {
	return []model.File{
		FileOne,
	}
}
