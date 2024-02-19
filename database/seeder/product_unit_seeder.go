package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	ProductUnitOne = model.ProductUnit{
		Id:          "cec71c3c-5a20-461b-9370-051ab3eeba76",
		ToUnitId:    nil,
		ImageFileId: nil,
		UnitId:      UnitKaleng.Id,
		ProductId:   ProductKalengCatMerah.Id,
		Scale:       1,
		ScaleToBase: 1,
	}
)

func ProductUnitSeeder(repositoryManager repository.RepositoryManager) {
	productUnitRepository := repositoryManager.ProductUnitRepository()

	count, err := productUnitRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := productUnitRepository.InsertMany(context.Background(), getProductUnitData()); err != nil {
		panic(err)
	}
}

func getProductUnitData() []model.ProductUnit {
	return []model.ProductUnit{
		ProductUnitOne,
	}
}
