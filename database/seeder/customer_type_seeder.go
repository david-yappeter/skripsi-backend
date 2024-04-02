package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	CustomerTypeOne = model.CustomerType{
		Id:          "717244a6-8318-49ca-b61b-f8cb1e58a63d",
		Name:        "Langganan Tier I",
		Description: nil,
	}
)

func CustomerTypeSeeder(repositoryManager repository.RepositoryManager) {
	customerTypeRepository := repositoryManager.CustomerTypeRepository()

	count, err := customerTypeRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := customerTypeRepository.InsertMany(context.Background(), getCustomerTypeData()); err != nil {
		panic(err)
	}
}

func getCustomerTypeData() []model.CustomerType {
	return []model.CustomerType{
		CustomerTypeOne,
	}
}
