package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/repository"
)

type ShopOrderItemUseCase interface {
	// webhook
	WebhookOrderStatusChange(ctx context.Context, request dto_request.TiktokWebhookBaseRequest[dto_request.WebhookOrderStatusChangeRequest])
}

type shopOrderItemUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewShopOrderItemUseCase(
	repositoryManager repository.RepositoryManager,
) ShopOrderItemUseCase {
	return &shopOrderItemUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *shopOrderItemUseCase) WebhookOrderStatusChange(ctx context.Context, request dto_request.TiktokWebhookBaseRequest[dto_request.WebhookOrderStatusChangeRequest]) {
	
}
