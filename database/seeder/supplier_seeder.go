package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

var (
	SupplierOne = model.Supplier{
		Id:             util.NewUuid(),
		SupplierTypeId: SupplierTypeOne.Id,
		Code:           "S-123",
		Name:           "Supplier A",
		IsActive:       true,
		Address:        "Jln. Tilak",
		Phone:          "+6285286869797",
		Email:          nil,
		Description:    nil,
	}
)

func SupplierSeeder(repositoryManager repository.RepositoryManager) {
	supplierRepository := repositoryManager.SupplierRepository()

	count, err := supplierRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := supplierRepository.InsertMany(context.Background(), getSupplierData()); err != nil {
		panic(err)
	}
}

func getSupplierData() []model.Supplier {
	return []model.Supplier{
		SupplierOne,
	}
}
