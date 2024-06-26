package use_case

import (
	"context"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/sync/errgroup"
)

type userLoaderParams struct {
	userRoles bool
}

type UserUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.UserCreateRequest) model.User
	AddRole(ctx context.Context, request dto_request.UserAddRoleRequest) model.User

	// read
	Fetch(ctx context.Context, request dto_request.UserFetchRequest) ([]model.User, int)
	Get(ctx context.Context, request dto_request.UserGetRequest) model.User
	GetMe(ctx context.Context) model.User

	//  update
	Update(ctx context.Context, request dto_request.UserUpdateRequest) model.User
	UpdatePassword(ctx context.Context, request dto_request.UserUpdatePasswordRequest) model.User

	//  delete
	DeleteRole(ctx context.Context, request dto_request.UserDeleteRoleRequest) model.User

	// option
	OptionForCashierSessionFilter(ctx context.Context, request dto_request.UserOptionForCashierSessionFilterRequest) ([]model.User, int)
	OptionForDeliveryOrderDriverForm(ctx context.Context, request dto_request.UserOptionForDeliveryOrderDriverFormRequest) ([]model.User, int)
	OptionForProductStockAdjustmentFilter(ctx context.Context, request dto_request.UserOptionForProductStockAdjustmentFilterRequest) ([]model.User, int)
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
		panic(dto_response.NewBadRequestErrorResponse("USER.UNIQUE_USERNAME"))
	}
}

func (u *userUseCase) mustLoadUsersData(ctx context.Context, users []*model.User, option userLoaderParams) {
	userRoleLoader := loader.NewUserRolesLoader(u.repositoryManager.UserRoleRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range users {
				if option.userRoles {
					group.Go(userRoleLoader.UserFn(users[i]))
				}
			}
		}),
	)

	roleLoader := loader.NewRoleLoader(u.repositoryManager.RoleRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range users {
				for j := range users[i].UserRoles {
					if option.userRoles {
						group.Go(roleLoader.UserRoleFn(&users[i].UserRoles[j]))
					}
				}
			}
		}),
	)
}

func (u *userUseCase) Create(ctx context.Context, request dto_request.UserCreateRequest) model.User {
	u.mustValidateUsernameUnique(ctx, request.Username)

	user := model.User{
		Id:       util.NewUuid(),
		Username: request.Username,
		Name:     request.Name,
		Password: u.mustGetHashedPassword(request.Password),
		IsActive: request.IsActive,
	}

	panicIfErr(
		u.repositoryManager.UserRepository().Insert(ctx, &user),
	)

	return user
}

func (u *userUseCase) AddRole(ctx context.Context, request dto_request.UserAddRoleRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)
	mustGetRole(ctx, u.repositoryManager, request.RoleId, true)

	userRole, err := u.repositoryManager.UserRoleRepository().GetByUserIdAndRoleId(ctx, request.UserId, request.RoleId)
	panicIfErr(err, constant.ErrNoData)

	if userRole != nil {
		panic(dto_response.NewBadRequestErrorResponse("USER.ROLE_ALREADY_EXIST"))
	}

	userRole = &model.UserRole{
		UserId: request.UserId,
		RoleId: request.RoleId,
	}

	panicIfErr(
		u.repositoryManager.UserRoleRepository().InsertMany(ctx, []model.UserRole{*userRole}),
	)

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) Fetch(ctx context.Context, request dto_request.UserFetchRequest) ([]model.User, int) {
	queryOption := model.UserQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase:   request.Phrase,
		IsActive: request.IsActive,
		RoleIds:  request.RoleIds,
	}

	users, err := u.repositoryManager.UserRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.UserRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadUsersData(ctx, util.SliceValueToSlicePointer(users), userLoaderParams{
		userRoles: true,
	})

	return users, total
}

func (u *userUseCase) Get(ctx context.Context, request dto_request.UserGetRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) Update(ctx context.Context, request dto_request.UserUpdateRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)

	user.Name = request.Name
	user.IsActive = request.IsActive

	panicIfErr(
		u.repositoryManager.UserRepository().Update(ctx, &user),
	)

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) UpdatePassword(ctx context.Context, request dto_request.UserUpdatePasswordRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)

	// change user password
	user.Password = u.mustGetHashedPassword(request.Password)

	// update all updated user non-revoked access token to revoked
	userAccessTokens, err := u.repositoryManager.UserAccessTokenRepository().FetchByUserIdAndRevokedAndGetExpiredAt(ctx, request.UserId, false, util.CurrentDateTime())
	panicIfErr(err)

	toBeRevokedUserAccessTokenIds := []string{}
	for _, userAccessToken := range userAccessTokens {
		toBeRevokedUserAccessTokenIds = append(toBeRevokedUserAccessTokenIds, userAccessToken.Id)
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			userRepository := u.repositoryManager.UserRepository()
			userAccessTokenRepository := u.repositoryManager.UserAccessTokenRepository()

			// update password
			if err := userRepository.Update(ctx, &user); err != nil {
				return err
			}

			// revoked access token
			if err := userAccessTokenRepository.UpdateManyRevokedByIds(ctx, true, toBeRevokedUserAccessTokenIds); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) DeleteRole(ctx context.Context, request dto_request.UserDeleteRoleRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)
	mustGetRole(ctx, u.repositoryManager, request.RoleId, false)
	userRole := mustGetUserRoleByUserIdAndRoleId(ctx, u.repositoryManager, request.UserId, request.RoleId, true)

	panicIfErr(
		u.repositoryManager.UserRoleRepository().Delete(ctx, &userRole),
	)

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) GetMe(ctx context.Context) model.User {
	authUser := model.MustGetUserCtx(ctx)

	u.mustLoadUsersData(ctx, []*model.User{&authUser}, userLoaderParams{
		userRoles: true,
	})

	userRoleIds := []string{}
	for _, role := range authUser.Roles {
		userRoleIds = append(userRoleIds, role.Id)
	}

	authUser.Permissions = []model.Permission{}

	permissions, err := u.repositoryManager.PermissionRepository().FetchByRoleIds(ctx, userRoleIds)
	panicIfErr(err)

	authUser.Permissions = permissions

	return authUser
}

func (u *userUseCase) OptionForCashierSessionFilter(ctx context.Context, request dto_request.UserOptionForCashierSessionFilterRequest) ([]model.User, int) {
	queryOption := model.UserQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts{{Field: "name", Direction: "asc"}},
		),
		Phrase: request.Phrase,
	}

	users, err := u.repositoryManager.UserRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.UserRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return users, total
}

func (u *userUseCase) OptionForDeliveryOrderDriverForm(ctx context.Context, request dto_request.UserOptionForDeliveryOrderDriverFormRequest) ([]model.User, int) {
	role, err := u.repositoryManager.RoleRepository().GetByName(ctx, data_type.RoleDriver)
	panicIfErr(err)

	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)
	deliveryOrderDriversLoader := loader.NewDeliveryOrderDriversLoader(u.repositoryManager.DeliveryOrderDriverRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		group.Go(deliveryOrderDriversLoader.DeliveryOrderFn(&deliveryOrder))
	}))

	excludedUserIds := []string{}
	for _, deliveryOrderDriver := range deliveryOrder.DeliveryOrderDrivers {
		excludedUserIds = append(excludedUserIds, deliveryOrderDriver.DriverUserId)
	}

	queryOption := model.UserQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts{{Field: "name", Direction: "asc"}},
		),
		ExcludedIds: excludedUserIds,
		Phrase:      request.Phrase,
		RoleIds:     []string{role.Id},
	}

	users, err := u.repositoryManager.UserRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.UserRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return users, total
}

func (u *userUseCase) OptionForProductStockAdjustmentFilter(ctx context.Context, request dto_request.UserOptionForProductStockAdjustmentFilterRequest) ([]model.User, int) {
	queryOption := model.UserQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts{{Field: "name", Direction: "asc"}},
		),
		Phrase: request.Phrase,
	}

	users, err := u.repositoryManager.UserRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.UserRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return users, total
}
