package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

var (
	ProductKalengCatMerah = model.Product{
		Id:          "e1bf0592-7850-4602-a740-6aae98dfd281",
		Name:        "Kaleng Cat Merah",
		Description: nil,
		Price:       util.Float64P(150000),
		IsActive:    true,
	}
)

func ProductSeeder(repositoryManager repository.RepositoryManager) {
	productRepository := repositoryManager.ProductRepository()

	count, err := productRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := productRepository.InsertMany(context.Background(), getProductData()); err != nil {
		panic(err)
	}
}

func getProductData() []model.Product {
	return []model.Product{
		ProductKalengCatMerah,
	}
}
