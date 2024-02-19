package use_case

import (
	"context"
	"log"
	"myapp/repository"
)

type TiktokConfigUseCase interface {
	//  delete
	AutoUpdate(ctx context.Context)
}

type tiktokConfigUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewTiktokConfigUseCase(
	repositoryManager repository.RepositoryManager,
) TiktokConfigUseCase {
	return &tiktokConfigUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *tiktokConfigUseCase) AutoUpdate(ctx context.Context) {
	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.RefreshToken == nil {
		log.Println("TIKTOK_CONFIG.REFRESH_TOKEN_EMPTY")
	}

	resp, err := client.RefreshToken(ctx, *tiktokConfig.RefreshToken)
	if err != nil {
		log.Println(err)
	}

	tiktokConfig.AccessToken = &resp.AccessToken
	tiktokConfig.RefreshToken = &resp.RefreshToken

	panicIfErr(
		u.repositoryManager.TiktokConfigRepository().Update(ctx, &tiktokConfig),
	)
}
