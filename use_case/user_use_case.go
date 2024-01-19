package use_case

import (
	"context"
	"myapp/constant"
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

	//  update
	Update(ctx context.Context, request dto_request.UserUpdateRequest) model.User
	UpdatePassword(ctx context.Context, request dto_request.UserUpdatePasswordRequest) model.User
	UpdateActive(ctx context.Context, request dto_request.UserUpdateActiveRequest) model.User
	UpdateInActive(ctx context.Context, request dto_request.UserUpdateInActiveRequest) model.User

	//  delete
	DeleteRole(ctx context.Context, request dto_request.UserDeleteRoleRequest) model.User

	// read
	GetMe(ctx context.Context) model.User
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
		Name:     u.mustGetHashedPassword(request.Password),
		Password: request.Password,
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

func (u *userUseCase) Update(ctx context.Context, request dto_request.UserUpdateRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)

	if user.Name != request.Name {
		user.Name = request.Name

		panicIfErr(
			u.repositoryManager.UserRepository().UpdateName(ctx, &user),
		)
	}

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) UpdatePassword(ctx context.Context, request dto_request.UserUpdatePasswordRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)

	user.Password = u.mustGetHashedPassword(request.Password)
	panicIfErr(
		u.repositoryManager.UserRepository().UpdatePassword(ctx, &user),
	)

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) UpdateActive(ctx context.Context, request dto_request.UserUpdateActiveRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)

	if user.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("USER.ALREADY_ACTIVE"))
	}

	user.IsActive = true
	panicIfErr(
		u.repositoryManager.UserRepository().UpdateIsActive(ctx, &user),
	)

	u.mustLoadUsersData(ctx, []*model.User{&user}, userLoaderParams{
		userRoles: true,
	})

	return user
}

func (u *userUseCase) UpdateInActive(ctx context.Context, request dto_request.UserUpdateInActiveRequest) model.User {
	user := mustGetUser(ctx, u.repositoryManager, request.UserId, false)

	if !user.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("USER.ALREADY.INACTIVE"))
	}

	user.IsActive = false
	panicIfErr(
		u.repositoryManager.UserRepository().UpdateIsActive(ctx, &user),
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

	return authUser
}
