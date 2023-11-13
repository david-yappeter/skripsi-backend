package use_case

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	jwtInternal "myapp/internal/jwt"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	// create
	GenerateJWT(ctx context.Context, userId string) (*jwtInternal.Token, error)
	LoginUsername(ctx context.Context, request dto_request.AuthUsernameLoginRequest) model.Token
	RegisterUsername(ctx context.Context, request dto_request.AuthUsernameRegisterRequest) model.Token

	// update
	Logout(ctx context.Context)

	// read
	Parse(ctx context.Context, token string) (*model.UserAccessToken, *model.User, error)
}

type authUseCase struct {
	repositoryManager repository.RepositoryManager

	jwt jwtInternal.Jwt
}

func NewAuthUseCase(
	repositoryManager repository.RepositoryManager,
	jwt jwtInternal.Jwt,
) AuthUseCase {
	return &authUseCase{
		repositoryManager: repositoryManager,
		jwt:               jwt,
	}
}

func (u *authUseCase) mustGetHashedPassword(originalPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	panicIfErr(err)
	return string(hashedPassword)
}

func (u *authUseCase) mustValidateComparePassword(hashedPassword string, originalPassword string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(originalPassword))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			panic(dto_response.NewBadRequestErrorResponse("AUTH.WRONG_PASSWORD"))
		}
		panic(err)
	}
}

func (u *authUseCase) mustValidateUsernameUnique(ctx context.Context, username string) {
	isExist, err := u.repositoryManager.UserRepository().IsExistByUsername(ctx, username)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("AUTH.UNIQUE_USERNAME"))
	}
}

func (u *authUseCase) generateAccessTokenId() (string, error) {
	bytes := make([]byte, 40)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", bytes), nil
}

func (u *authUseCase) generateUserAccessToken(ctx context.Context, userId string) (*model.UserAccessToken, error) {
	var (
		now                                  = util.CurrentDateTime()
		duration                             = time.Hour * 24
		expiredAt                            = now.Add(duration)
		maxAccessTokenGenerationAttempts int = 10
	)

	userAccessToken := &model.UserAccessToken{
		UserId:    userId,
		Revoked:   false,
		ExpiredAt: expiredAt,
	}

	for maxAccessTokenGenerationAttempts > 0 {
		maxAccessTokenGenerationAttempts--

		accessTokenId, err := u.generateAccessTokenId()
		if err != nil {
			return nil, err
		}

		userAccessToken.Id = accessTokenId

		if err := u.repositoryManager.UserAccessTokenRepository().Insert(ctx, userAccessToken); err != nil && maxAccessTokenGenerationAttempts == 0 {
			log.Println(err)
			return nil, errors.New("max access token generation attempts exceeded")
		} else if err == nil {
			break
		}
	}

	return userAccessToken, nil
}

func (u *authUseCase) GenerateJWT(ctx context.Context, userId string) (*jwtInternal.Token, error) {
	userAccessToken, err := u.generateUserAccessToken(ctx, userId)
	if err != nil {
		return nil, err
	}

	accessToken, err := u.jwt.Generate(jwtInternal.Payload{
		Id:        userAccessToken.Id,
		UserId:    userAccessToken.UserId,
		CreatedAt: userAccessToken.CreatedAt.Time(),
		ExpiredAt: userAccessToken.ExpiredAt.Time(),
	})
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (u *authUseCase) LoginUsername(ctx context.Context, request dto_request.AuthUsernameLoginRequest) model.Token {
	user, err := u.repositoryManager.UserRepository().GetByUsernameAndIsActive(ctx, request.Username, true)
	panicIfErr(err, constant.ErrNoData)
	if user == nil {
		panic(dto_response.NewForbiddenErrorResponse("Sorry, your account not registered in our system"))
	}

	u.mustValidateComparePassword(user.Password, request.Password)

	accessToken, err := u.GenerateJWT(ctx, user.Id)
	panicIfErr(err)

	return model.Token{
		AccessToken:          accessToken.AccessToken,
		AccessTokenExpiredAt: data_type.NewDateTime(accessToken.ExpiredAt),
		TokenType:            accessToken.TokenType,
	}
}

func (u *authUseCase) RegisterUsername(ctx context.Context, request dto_request.AuthUsernameRegisterRequest) model.Token {
	u.mustValidateUsernameUnique(ctx, request.Username)

	hashedPassword := u.mustGetHashedPassword(request.Password)

	user := model.User{
		Id:       util.NewUuid(),
		Username: request.Username,
		Name:     request.Name,
		Password: hashedPassword,
		IsActive: request.IsActive,
	}

	panicIfErr(
		u.repositoryManager.UserRepository().Insert(ctx, &user),
	)

	accessToken, err := u.GenerateJWT(ctx, user.Id)
	panicIfErr(err)

	return model.Token{
		AccessToken:          accessToken.AccessToken,
		AccessTokenExpiredAt: data_type.NewDateTime(accessToken.ExpiredAt),
		TokenType:            accessToken.TokenType,
	}
}

func (u *authUseCase) Logout(ctx context.Context) {
	userAccessToken := model.MustGetUserAccessTokenCtx(ctx)

	userAccessToken.Revoked = true

	panicIfErr(
		u.repositoryManager.UserAccessTokenRepository().UpdateRevoked(ctx, &userAccessToken),
	)
}

func (u *authUseCase) Parse(ctx context.Context, token string) (*model.UserAccessToken, *model.User, error) {
	payload, err := u.jwt.Parse(token)
	if err != nil {
		return nil, nil, constant.ErrNotAuthenticated
	}

	var (
		accessTokenId = payload.Id
		userId        = payload.UserId
	)

	userAccessToken, err := u.repositoryManager.UserAccessTokenRepository().Get(ctx, accessTokenId)
	if err != nil {
		if err == constant.ErrNoData {
			return nil, nil, constant.ErrNotAuthenticated
		}

		return nil, nil, err
	}

	if userAccessToken.UserId != userId || userAccessToken.Revoked || userAccessToken.ExpiredAt.IsLessThan(util.CurrentDateTime()) {
		return nil, nil, constant.ErrNotAuthenticated
	}

	user, err := u.repositoryManager.UserRepository().GetByIdAndIsActive(ctx, userId, true)
	if err != nil {
		if err == constant.ErrNoData {
			return nil, nil, constant.ErrNotAuthenticated
		}

		return nil, nil, err
	}

	roles, err := u.repositoryManager.RoleRepository().FetchByUserId(ctx, user.Id)
	if err != nil {
		return nil, nil, err
	}

	user.Roles = roles

	// user.RoleTypes = map[data_type.RoleType]struct{}{}
	// for _, role := range roles {
	// 	user.RoleTypes[role.RoleType()] = struct{}{}
	// }

	// if user.HasDoctorRoleType() {
	// 	doctor, err := u.doctorRepository.GetByUserId(ctx, user.Id)
	// 	if err != nil && err != constant.ErrNoData {
	// 		return nil, nil, err
	// 	}

	// 	user.Doctor = doctor
	// }

	// if user.HasNurseRoleType() {
	// 	nurse, err := u.nurseRepository.GetByUserId(ctx, user.Id)
	// 	if err != nil && err != constant.ErrNoData {
	// 		return nil, nil, err
	// 	}

	// 	user.Nurse = nurse
	// }

	// if user.HasPharmacistRoleType() {
	// 	pharmacist, err := u.pharmacistRepository.GetByUserId(ctx, user.Id)
	// 	if err != nil && err != constant.ErrNoData {
	// 		return nil, nil, err
	// 	}

	// 	user.Pharmacist = pharmacist
	// }

	return userAccessToken, user, nil
}
