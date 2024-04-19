package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

var (
	ShopeeConfigSeederOne = model.ShopeeConfig{
		PartnerId:    "",
		PartnerKey:   "",
		AccessToken:  nil,
		RefreshToken: nil,
	}
)

func ShopeeConfigSeeder(repositoryManager repository.RepositoryManager) {
	shopeeConfigRepository := repositoryManager.ShopeeConfigRepository()

	count, err := shopeeConfigRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := shopeeConfigRepository.InsertMany(context.Background(), getShopeeConfigData()); err != nil {
		panic(err)
	}
}

func getShopeeConfigData() []model.ShopeeConfig {
	return []model.ShopeeConfig{
		ShopeeConfigSeederOne,
	}
}
