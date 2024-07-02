package use_case

import (
	"context"
	"fmt"
	"log"
	"math"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/infrastructure"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/repository"
	"myapp/util"
	"net/http"
	"strings"

	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"
)

type WhatsappUseCase interface {
	// utility
	IsLoggedIn(ctx context.Context) bool
	Login(ctx context.Context) chan (string)
	Logout(ctx context.Context)

	// broadcast
	CustomerDebtBroadcast(ctx context.Context, request dto_request.WhatsappCustomerDebtBroadcastRequest)
	CustomerTypeDiscountBroadcast(ctx context.Context, request dto_request.WhatsappCustomerTypeDiscountBroadcastRequest)
	ProductPriceChangeBroadcast(ctx context.Context, request dto_request.WhatsappProductPriceChangeBroadcastRequest)
}

type whatsappUseCase struct {
	repositoryManager repository.RepositoryManager
	whatsappManager   *infrastructure.WhatsappManager

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client
}

func NewWhatsappUseCase(
	repositoryManager repository.RepositoryManager,
	whatsappManager *infrastructure.WhatsappManager,
	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) WhatsappUseCase {
	return &whatsappUseCase{
		repositoryManager: repositoryManager,
		whatsappManager:   whatsappManager,

		mainFilesystem: mainFilesystem,
		tmpFilesystem:  tmpFilesystem,
	}
}

func (u *whatsappUseCase) IsLoggedIn(ctx context.Context) bool {
	if u.whatsappManager == nil {
		return false
	}

	return (*u.whatsappManager).IsLoggedIn(ctx)
}

func (u *whatsappUseCase) Login(ctx context.Context) chan (string) {
	if u.whatsappManager == nil {
		return nil
	}

	qrString, _ := (*u.whatsappManager).LoginQr(ctx)
	return qrString
}

func (u *whatsappUseCase) Logout(ctx context.Context) {
	if u.whatsappManager == nil {
		return
	}

	panicIfErr(
		(*u.whatsappManager).Logout(),
	)
}

func (u *whatsappUseCase) CustomerDebtBroadcast(ctx context.Context, request dto_request.WhatsappCustomerDebtBroadcastRequest) {
	if u.whatsappManager == nil {
		return
	}

	customer := mustGetCustomer(ctx, u.repositoryManager, request.CustomerId, true)

	customerDebts, err := u.repositoryManager.CustomerDebtRepository().FetchByCustomerIdAndStatuses(
		ctx,
		customer.Id,
		[]data_type.CustomerDebtStatus{
			data_type.CustomerDebtStatusUnpaid,
			data_type.CustomerDebtStatusHalfPaid,
		},
	)
	panicIfErr(err)

	deliveryOrderLoader := loader.NewDeliveryOrderLoader(u.repositoryManager.DeliveryOrderRepository())
	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range customerDebts {
			group.Go(deliveryOrderLoader.CustomerDebtFn(&customerDebts[i]))
		}
	}))

	totalRemainingDebts := 0.0
	customerDebtListMessage := ``

	for _, customerDebt := range customerDebts {
		totalRemainingDebts += customerDebt.RemainingAmount

		customerDebtListMessage += fmt.Sprintf(`- Pengiriman %s, Total yang belum dibayar: Rp. %s
`, customerDebt.DeliveryOrder.InvoiceNumber, util.CurrencyFormat(int(customerDebt.RemainingAmount), language.Indonesian))
	}

	messageTemplate := `Halo %s,

Kami berharap Anda dalam keadaan sehat dan baik.

Kami ingin mengingatkan Anda bahwa terdapat beberapa tagihan yang masih belum diselesaikan. Berikut adalah rincian tagihan Anda:

%s

Total Jumlah yang Harus Dibayar: Rp. *%s*

Jika Anda sudah melakukan pembayaran, silakan abaikan pesan ini atau hubungi kami untuk konfirmasi pembayaran. Jika ada pertanyaan atau membutuhkan bantuan lebih lanjut, jangan ragu untuk menghubungi kami di nomor ini.

Terima kasih atas perhatian dan kerjasamanya.

Hormat kami,
*%s*`

	go func() {
		goCtx := context.Background()

		customerJID, _ := types.ParseJID(fmt.Sprintf("%s@s.whatsapp.net", strings.Trim(customer.Phone, "+")))

		(*u.whatsappManager).SendMessage(goCtx, customerJID, &proto.Message{
			Conversation: util.Pointer(
				fmt.Sprintf(messageTemplate,
					customer.Name,
					customerDebtListMessage,
					util.CurrencyFormat(int(totalRemainingDebts), language.Indonesian),
					"Toko Setia Abadi",
				),
			),
		})
	}()
}

func (u *whatsappUseCase) CustomerTypeDiscountBroadcast(ctx context.Context, request dto_request.WhatsappCustomerTypeDiscountBroadcastRequest) {
	if u.whatsappManager == nil {
		return
	}

	customerTypeDiscount := mustGetCustomerTypeDiscount(ctx, u.repositoryManager, request.CustomerTypeDiscountId, true)
	product := mustGetProduct(ctx, u.repositoryManager, customerTypeDiscount.ProductId, true)

	if product.Price == nil {
		panic(dto_response.NewBadRequestErrorResponse("WHATSAPP.PRODUCT_MUST_HAVE_PRICE"))
	}

	discountAmount := 0.0
	if customerTypeDiscount.DiscountAmount != nil {
		discountAmount = *customerTypeDiscount.DiscountAmount
	} else {
		discountAmount = *customerTypeDiscount.DiscountPercentage * *product.Price / 100.0
	}

	imageFile := mustGetFile(ctx, u.repositoryManager, product.ImageFileId, true)

	// get image bytes
	data, err := u.mainFilesystem.ReadFile(imageFile.Path)
	panicIfErr(err)

	// guess mimetypes
	mimeType := http.DetectContentType(data)

	customers, err := u.repositoryManager.CustomerRepository().FetchByCustomerTypeId(ctx, &customerTypeDiscount.CustomerTypeId)
	panicIfErr(err)

	messageTemplate := `🌟 Diskon Khusus untuk Anda, Pelanggan Istimewa!

Halo %s,

Kami ingin mengucapkan terima kasih atas dukungan Anda sebagai pelanggan istimewa kami! Sebagai bentuk apresiasi, kami ingin menawarkan penawaran eksklusif berikut kepada Anda:

🎁 Diskon Khusus untuk Produk *%s*!

Sekarang adalah kesempatan Anda untuk mendapatkan produk favorit Anda dengan harga istimewa. Jangan lewatkan kesempatan ini!

💰Harga Lama			: ~Rp. %s~
💸Harga Setelah Diskon	: *Rp. %s*

Jika ada pertanyaan atau butuh bantuan, jangan ragu untuk menghubungi kami. Kami siap membantu Anda dengan senang hati.

Terima kasih atas kesetiaan dan dukungan Anda kepada kami!

Salam hangat,
*%s*`

	if len(customers) > 0 {
		go func() {
			goCtx := context.Background()
			// upload image to whatsapp
			resp, err := (*u.whatsappManager).UploadImage(goCtx, data)
			if err != nil {
				log.Println(err)
				return
			}

			// send message
			if resp != nil {
				for _, customer := range customers {
					customerJID, _ := types.ParseJID(fmt.Sprintf("%s@s.whatsapp.net", strings.Trim(customer.Phone, "+")))
					(*u.whatsappManager).SendMessage(goCtx, customerJID, &proto.Message{
						ImageMessage: &proto.ImageMessage{
							Caption: util.StringP(fmt.Sprintf(messageTemplate,
								customer.Name,
								product.Name,
								util.CurrencyFormat(int(*product.Price), language.Indonesian),
								util.CurrencyFormat(int(math.Max(0, *product.Price-discountAmount)), language.Indonesian),
								"Toko Setia Abadi",
							)),
							Mimetype: util.StringP(mimeType),

							Url:           &resp.URL,
							DirectPath:    &resp.DirectPath,
							MediaKey:      resp.MediaKey,
							FileEncSha256: resp.FileEncSHA256,
							FileSha256:    resp.FileSHA256,
							FileLength:    &resp.FileLength,
						},
					})
				}
			}
		}()
	}
}

func (u *whatsappUseCase) ProductPriceChangeBroadcast(ctx context.Context, request dto_request.WhatsappProductPriceChangeBroadcastRequest) {
	if u.whatsappManager == nil {
		return
	}

	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	if product.Price == nil {
		panic(dto_response.NewBadRequestErrorResponse("WHATSAPP.PRODUCT_MUST_HAVE_PRICE"))
	}

	imageFile := mustGetFile(ctx, u.repositoryManager, product.ImageFileId, true)

	// get image bytes
	data, err := u.mainFilesystem.ReadFile(imageFile.Path)
	panicIfErr(err)

	// guess mimetypes
	mimeType := http.DetectContentType(data)

	customers, err := u.repositoryManager.CustomerRepository().FetchByCustomerTypeId(ctx, &request.CustomerTypeId)
	panicIfErr(err)

	messageTemplate := `🛍️ Pemberitahuan Pergantian Harga Barang

Halo %s,

Kami ingin memberitahu Anda tentang perubahan harga pada salah satu barang yang Anda minati. Dalam rangka memastikan ketersediaan stok dan kualitas layanan kami, harga untuk barang berikut telah mengalami perubahan:

📦 Nama Barang	: *%s*
💰 Harga Lama	: ~Rp. %s~
💸 Harga Baru	: *Rp. %s*

Kami memahami bahwa perubahan harga mungkin dapat mempengaruhi keputusan Anda. Namun, kami berkomitmen untuk terus memberikan produk berkualitas dengan harga yang kompetitif.

Jika Anda memiliki pertanyaan atau membutuhkan klarifikasi lebih lanjut, jangan ragu untuk menghubungi kami. Kami siap membantu Anda dengan setulus hati.

Terima kasih atas pengertian dan dukungan Anda.

Salam hangat,
*%s*`

	if len(customers) > 0 {
		go func() {
			goCtx := context.Background()
			// upload image to whatsapp
			resp, err := (*u.whatsappManager).UploadImage(goCtx, data)
			if err != nil {
				log.Println(err)
				return
			}

			// send message
			if resp != nil {
				for _, customer := range customers {
					customerJID, _ := types.ParseJID(fmt.Sprintf("%s@s.whatsapp.net", strings.Trim(customer.Phone, "+")))
					(*u.whatsappManager).SendMessage(goCtx, customerJID, &proto.Message{
						ImageMessage: &proto.ImageMessage{
							Caption: util.StringP(fmt.Sprintf(messageTemplate,
								customer.Name,
								product.Name,
								util.CurrencyFormat(int(request.OldPrice), language.Indonesian),
								util.CurrencyFormat(int(*product.Price), language.Indonesian),
								"Toko Setia Abadi",
							)),
							Mimetype: util.StringP(mimeType),

							Url:           &resp.URL,
							DirectPath:    &resp.DirectPath,
							MediaKey:      resp.MediaKey,
							FileEncSha256: resp.FileEncSHA256,
							FileSha256:    resp.FileSHA256,
							FileLength:    &resp.FileLength,
						},
					})
				}
			}
		}()
	}
}
