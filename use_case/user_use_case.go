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
	// admin create
	AdminCreate(ctx context.Context, request dto_request.AdminUserCreateRequest) model.User
	AdminAddRole(ctx context.Context, request dto_request.AdminUserAddRoleRequest) model.User

	// admin update
	AdminUpdate(ctx context.Context, request dto_request.AdminUserUpdateRequest) model.User
	AdminUpdatePassword(ctx context.Context, request dto_request.AdminUserUpdatePasswordRequest) model.User
	AdminUpdateActive(ctx context.Context, request dto_request.AdminUserUpdateActiveRequest) model.User
	AdminUpdateInActive(ctx context.Context, request dto_request.AdminUserUpdateInActiveRequest) model.User

	// admin delete
	AdminDeleteRole(ctx context.Context, request dto_request.AdminUserDeleteRoleRequest) model.User

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
				for j := range users[i].Roles {
					if option.userRoles {
						group.Go(roleLoader.UserRoleFn(&users[i].UserRoles[j]))
					}
				}
			}
		}),
	)
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

func (u *userUseCase) AdminAddRole(ctx context.Context, request dto_request.AdminUserAddRoleRequest) model.User {
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

func (u *userUseCase) AdminUpdate(ctx context.Context, request dto_request.AdminUserUpdateRequest) model.User {
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

func (u *userUseCase) AdminUpdatePassword(ctx context.Context, request dto_request.AdminUserUpdatePasswordRequest) model.User {
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

func (u *userUseCase) AdminUpdateActive(ctx context.Context, request dto_request.AdminUserUpdateActiveRequest) model.User {
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

func (u *userUseCase) AdminUpdateInActive(ctx context.Context, request dto_request.AdminUserUpdateInActiveRequest) model.User {
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

func (u *userUseCase) AdminDeleteRole(ctx context.Context, request dto_request.AdminUserDeleteRoleRequest) model.User {
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
