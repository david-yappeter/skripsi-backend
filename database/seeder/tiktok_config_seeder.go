package seeder

import (
	"context"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

var (
	TiktokConfigSeederOne = model.TiktokConfig{
		AppKey:       "6bgdj91pdnm9v",
		AppSecret:    "6cfbf8374b80618ca0d1b5eafd87ca0e23554e57",
		WarehouseId:  "7333151372009178885",
		ShopId:       "7495591168837323491",
		ShopCipher:   "ROW_ij-EHgAAAAAFH7_LWApa2DADTZh6ANIA",
		AccessToken:  nil,
		RefreshToken: util.StringP("ROW_FAdt_wAAAACx9YYnlqBwGkIEDGqW7sd7PMETXQn14yBwu76XelX093IZsWgM5bXG6Sx2rbflZ7A"),
		Timestamp:    model.Timestamp{},
	}
)

func TiktokConfigSeeder(repositoryManager repository.RepositoryManager) {
	tiktokConfigRepository := repositoryManager.TiktokConfigRepository()

	count, err := tiktokConfigRepository.Count(context.Background())
	if err != nil {
		panic(err)
	}

	// Stop if table already have data
	if count > 0 {
		return
	}

	if err := tiktokConfigRepository.InsertMany(context.Background(), getTiktokConfigData()); err != nil {
		panic(err)
	}
}

func getTiktokConfigData() []model.TiktokConfig {
	return []model.TiktokConfig{
		TiktokConfigSeederOne,
	}
}
