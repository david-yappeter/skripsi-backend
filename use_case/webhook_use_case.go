package use_case

import (
	"context"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	gotiktok "github.com/david-yappeter/go-tiktok"
)

type WebhookUseCase interface {
	OrderStatusChange(ctx context.Context, request dto_request.TiktokWebhookBaseRequest[dto_request.WebhookOrderStatusChangeRequest])
}

type webhookUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewWebhookUseCase(
	repositoryManager repository.RepositoryManager,
) WebhookUseCase {
	return &webhookUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *webhookUseCase) OrderStatusChange(ctx context.Context, request dto_request.TiktokWebhookBaseRequest[dto_request.WebhookOrderStatusChangeRequest]) {
	execWebhookMutex(func() {

		if request.Type != 1 {
			return
		}

		client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

		if tiktokConfig.AccessToken == nil {
			panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
		}

		orderList, err := client.GetOrderDetail(
			ctx,
			gotiktok.CommonParam{
				AccessToken: *tiktokConfig.AccessToken,
				ShopCipher:  tiktokConfig.ShopCipher,
				ShopId:      tiktokConfig.ShopId,
			},
			[]string{request.Data.OrderId},
		)
		panicIfErr(err)

		orderDetail := orderList.Orders[0]

		shopOrderItems := []model.ShopOrderItem{}

		shopOrder, err := u.repositoryManager.ShopOrderRepository().GetByPlatformTypeAndPlatformIdentifierId(ctx, data_type.ShopOrderPlatformTypeTiktokShop, orderDetail.Id)
		panicIfErr(err, constant.ErrNoData)

		isNewOrderData := shopOrder == nil

		if !isNewOrderData {
			var trackingNumber *string = nil
			if orderDetail.TrackingNumber != "" {
				trackingNumber = &orderDetail.TrackingNumber
			}

			nextStatus := data_type.DetermineOrderTrackingStatusByString(orderDetail.Status)
			if shopOrder.TrackingStatus.IsNextStatusValid(nextStatus) {
				shopOrder.TrackingStatus = nextStatus
			}
			shopOrder.TrackingNumber = trackingNumber
			shopOrder.RecipientName = orderDetail.RecipientAddress.Name
			shopOrder.RecipientFullAddress = orderDetail.RecipientAddress.FullAddress
			shopOrder.RecipientPhoneNumber = orderDetail.RecipientAddress.PhoneNumber
		} else {
			var trackingNumber *string = nil
			if orderDetail.TrackingNumber != "" {
				trackingNumber = &orderDetail.TrackingNumber
			}

			shopOrder = &model.ShopOrder{
				Id:                        util.NewUuid(),
				TrackingNumber:            trackingNumber,
				PlatformIdentifier:        orderDetail.Id,
				PlatformType:              data_type.ShopOrderPlatformTypeTiktokShop,
				TrackingStatus:            data_type.DetermineOrderTrackingStatusByString(orderDetail.Status),
				RecipientName:             orderDetail.RecipientAddress.Name,
				RecipientFullAddress:      orderDetail.RecipientAddress.FullAddress,
				RecipientPhoneNumber:      orderDetail.RecipientAddress.PhoneNumber,
				ShippingFee:               util.MustParseFloat64(orderDetail.Payment.ShippingFee),
				TotalOriginalProductPrice: util.MustParseFloat64(orderDetail.Payment.OriginalTotalProductPrice),
				Subtotal:                  util.MustParseFloat64(orderDetail.Payment.SubTotal),
				Tax:                       util.MustParseFloat64(orderDetail.Payment.Tax),
				TotalAmount:               util.MustParseFloat64(orderDetail.Payment.TotalAmount),
			}

			mapShopOrderItemByPlatformProductId := map[string]*model.ShopOrderItem{}

			for _, item := range orderDetail.LineItems {
				if mapShopOrderItemByPlatformProductId[item.ProductId] == nil {
					productUnit, err := u.repositoryManager.ProductUnitRepository().GetBaseProductUnitByProductId(ctx, item.SellerSku)
					panicIfErr(err)

					mapShopOrderItemByPlatformProductId[item.ProductId] = &model.ShopOrderItem{
						Id:                util.NewUuid(),
						ShopOrderId:       shopOrder.Id,
						ProductUnitId:     productUnit.Id,
						PlatformProductId: item.ProductId,
						ImageLink:         &item.SkuImage,
						Quantity:          1,
						OriginalPrice:     util.MustParseFloat64(item.OriginalPrice),
						SalePrice:         util.MustParseFloat64(item.SalePrice),
						Timestamp:         model.Timestamp{},
					}
				} else {
					mapShopOrderItemByPlatformProductId[item.ProductId].Quantity += 1
				}
			}
			for _, shopOrderItem := range mapShopOrderItemByPlatformProductId {
				shopOrderItems = append(shopOrderItems, *shopOrderItem)
			}
		}

		panicIfErr(
			u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
				shopOrderRepository := u.repositoryManager.ShopOrderRepository()
				shopOrderItemRepository := u.repositoryManager.ShopOrderItemRepository()

				if isNewOrderData {
					if err := shopOrderRepository.Insert(ctx, shopOrder); err != nil {
						return err
					}

					if err := shopOrderItemRepository.InsertMany(ctx, shopOrderItems); err != nil {
						return err
					}
				} else {
					if err := shopOrderRepository.Update(ctx, shopOrder); err != nil {
						return err
					}
				}

				return nil
			}),
		)
	})
}
