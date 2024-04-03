package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

var (
	SupplierTypeOne = model.SupplierType{
		Id:          util.NewUuid(),
		Name:        "Supplier A",
		Description: nil,
	}
)

func SupplierTypeSeeder(repositoryManager repository.RepositoryManager) {
	supplierTypeRepository := repositoryManager.SupplierTypeRepository()

	count, err := supplierTypeRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := supplierTypeRepository.InsertMany(context.Background(), getSupplierTypeData()); err != nil {
		panic(err)
	}
}

func getSupplierTypeData() []model.SupplierType {
	return []model.SupplierType{
		SupplierTypeOne,
	}
}
