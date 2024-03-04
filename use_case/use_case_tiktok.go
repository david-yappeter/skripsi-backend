package use_case

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	gotiktok "github.com/david-yappeter/go-tiktok"
)

func mustGetTiktokConfig(ctx context.Context, repositoryManager repository.RepositoryManager) model.TiktokConfig {
	tiktokConfig, err := repositoryManager.TiktokConfigRepository().Get(ctx)
	panicIfErr(err, constant.ErrNoData)

	if tiktokConfig == nil {
		panic("TIKTOK_CONFIG.NOT_SET")
	}

	return *tiktokConfig
}

func mustGetTiktokClient(ctx context.Context, repositoryManager repository.RepositoryManager) (*gotiktok.Client, model.TiktokConfig) {
	tiktokConfig := mustGetTiktokConfig(ctx, repositoryManager)

	client, err := gotiktok.New(
		tiktokConfig.AppKey,
		tiktokConfig.AppSecret,
		"202309",
	)
	panicIfErr(err)
	return client, tiktokConfig
}

func mustGetTiktokProductDetail(ctx context.Context, repositoryManager repository.RepositoryManager, tiktokProductId string) gotiktok.ProductDetailData {

	client, tiktokConfig := mustGetTiktokClient(ctx, repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	resp, err := client.GetProductDetail(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		tiktokProductId,
	)
	panicIfErr(err)

	return resp
}

func mustUpdateTiktokProductInventory(ctx context.Context, repositoryManager repository.RepositoryManager, tiktokProductId string, stockQtyCount int) {
	client, tiktokConfig := mustGetTiktokClient(ctx, repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	tiktokProductDetail := mustGetTiktokProductDetail(ctx, repositoryManager, tiktokProductId)

	_, err := client.UpdateProductInventory(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		tiktokProductId,
		gotiktok.UpdateProductInventoryRequest{
			Skus: []gotiktok.UpdateProductInventoryRequestSku{
				{
					Id: tiktokProductDetail.Skus[0].Id,
					Inventory: []gotiktok.UpdateProductInventoryRequestSkuInventory{
						{
							WarehouseId: tiktokProductDetail.Skus[0].Inventory[0].WarehouseId,
							Quantity:    int(stockQtyCount),
						},
					},
				},
			},
		},
	)
	panicIfErr(err)
}
