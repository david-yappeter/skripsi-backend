package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	UnitKaleng = model.Unit{
		Id:          "6745548c-ea48-4db8-b7d9-ed2cf1175ade",
		Name:        "Kaleng",
		Description: nil,
	}
	UnitDus = model.Unit{
		Id:          "3867eb2b-8905-402c-bce3-c5953262ec03",
		Name:        "Dus",
		Description: nil,
	}
)

func UnitSeeder(repositoryManager repository.RepositoryManager) {
	unitRepository := repositoryManager.UnitRepository()

	count, err := unitRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := unitRepository.InsertMany(context.Background(), getUnitData()); err != nil {
		panic(err)
	}
}

func getUnitData() []model.Unit {
	return []model.Unit{
		UnitKaleng,
		UnitDus,
	}
}
