package use_case

import (
	"context"
	"myapp/infrastructure"
	"myapp/repository"
)

type WhatsappUseCase interface {
	IsLoggedIn(ctx context.Context) bool
	Login(ctx context.Context) chan (string)
	Logout(ctx context.Context)
}

type whatsappUseCase struct {
	repositoryManager repository.RepositoryManager
	whatsappManager   *infrastructure.WhatsappManager
}

func NewWhatsappUseCase(
	repositoryManager repository.RepositoryManager,
	whatsappManager *infrastructure.WhatsappManager,
) WhatsappUseCase {
	return &whatsappUseCase{
		repositoryManager: repositoryManager,
		whatsappManager:   whatsappManager,
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
