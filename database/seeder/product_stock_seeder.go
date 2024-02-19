package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	ProductStockOne = model.ProductStock{
		Id:        "6feb27eb-d326-4234-ae0b-db810680fa57",
		ProductId: ProductKalengCatMerah.Id,
		Qty:       100,
	}
)

func ProductStockSeeder(repositoryManager repository.RepositoryManager) {
	productStockRepository := repositoryManager.ProductStockRepository()

	count, err := productStockRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := productStockRepository.InsertMany(context.Background(), getProductStockData()); err != nil {
		panic(err)
	}
}

func getProductStockData() []model.ProductStock {
	return []model.ProductStock{
		ProductStockOne,
	}
}
