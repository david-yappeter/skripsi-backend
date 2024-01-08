package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	// admin create
	AdminCreate(ctx context.Context, request dto_request.AdminUserCreateRequest) model.User

	// admin update
	AdminUpdate(ctx context.Context, request dto_request.AdminUserUpdateRequest) model.User
	AdminUpdatePassword(ctx context.Context, request dto_request.AdminUserUpdatePasswordRequest) model.User
	AdminUpdateActive(ctx context.Context, request dto_request.AdminUserUpdateActiveRequest) model.User
	AdminUpdateInActive(ctx context.Context, request dto_request.AdminUserUpdateInActiveRequest) model.User
}

type userUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewUserUseCase(
	repositoryManager repository.RepositoryManager,
) UserUseCase {
	return &userUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *userUseCase) mustGetHashedPassword(originalPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	panicIfErr(err)
	return string(hashedPassword)
}

func (u *userUseCase) mustValidateUsernameUnique(ctx context.Context, username string) {
	isExist, err := u.repositoryManager.UserRepository().IsExistByUsername(ctx, username)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("AUTH.UNIQUE_USERNAME"))
	}
}

func (u *userUseCase) AdminCreate(ctx context.Context, request dto_request.AdminUserCreateRequest) model.User {
	u.mustValidateUsernameUnique(ctx, request.Username)

	user := model.User{
		Id:       util.NewUuid(),
		Username: request.Username,
		Name:     u.mustGetHashedPassword(request.Password),
		Password: request.Password,
		IsActive: request.IsActive,
	}

	panicIfErr(
		u.repositoryManager.UserRepository().Insert(ctx, &user),
	)

	return user
}

func (u *userUseCase) AdminUpdate(ctx context.Context, request dto_request.AdminUserUpdateRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.Id, false)

	if user.Name != request.Name {
		user.Name = request.Name

		panicIfErr(
			u.repositoryManager.UserRepository().UpdateName(ctx, &user),
		)
	}

	return user
}

func (u *userUseCase) AdminUpdatePassword(ctx context.Context, request dto_request.AdminUserUpdatePasswordRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.Id, false)

	user.Password = u.mustGetHashedPassword(request.Password)
	panicIfErr(
		u.repositoryManager.UserRepository().UpdatePassword(ctx, &user),
	)

	return user
}

func (u *userUseCase) AdminUpdateActive(ctx context.Context, request dto_request.AdminUserUpdateActiveRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.Id, false)

	if user.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("USER.ALREADY_ACTIVE"))
	}

	user.IsActive = true
	panicIfErr(
		u.repositoryManager.UserRepository().UpdateIsActive(ctx, &user),
	)

	return user
}

func (u *userUseCase) AdminUpdateInActive(ctx context.Context, request dto_request.AdminUserUpdateInActiveRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.Id, false)

	if !user.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("USER.ALREADY.INACTIVE"))
	}

	user.IsActive = false
	panicIfErr(
		u.repositoryManager.UserRepository().UpdateIsActive(ctx, &user),
	)

	return user
}
