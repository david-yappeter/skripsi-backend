package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	CustomerOne = model.Customer{
		Id:             "717244a6-8318-49ca-b61b-f8cb1e58a63d",
		CustomerTypeId: &CustomerTypeOne.Id,
		Name:           "Bobby Doe",
		Email:          "bobby@gmail.com",
		Address:        "Jln. Mahkamah",
		Latitude:       3.574727,
		Longitude:      98.688233,
		Phone:          "+6285206069595",
		IsActive:       true,
	}
)

func CustomerSeeder(repositoryManager repository.RepositoryManager) {
	customerRepository := repositoryManager.CustomerRepository()

	count, err := customerRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := customerRepository.InsertMany(context.Background(), getCustomerData()); err != nil {
		panic(err)
	}
}

func getCustomerData() []model.Customer {
	return []model.Customer{
		CustomerOne,
	}
}
