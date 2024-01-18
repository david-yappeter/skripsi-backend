package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"mime/multipart"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_response"
	"myapp/delivery/middleware"
	"myapp/global"
	"myapp/manager"
	"myapp/model"
	"myapp/use_case"
	"myapp/util"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"

	filesystemInternal "myapp/internal/filesystem"
	bindingInternal "myapp/internal/gin/binding"
	"myapp/internal/gin/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
)

type apiContext struct {
	ginCtx *gin.Context
}

func newApiContext(ctx *gin.Context) apiContext {
	return apiContext{
		ginCtx: ctx,
	}
}

func (a *apiContext) context() context.Context {
	return a.ginCtx.Request.Context()
}

func (a *apiContext) getClientIp() string {
	u := &url.URL{Host: a.ginCtx.ClientIP()}
	return u.Hostname()
}

func (a *apiContext) getFile(key string) *multipart.FileHeader {
	// TODO: check MaxMultipartMemory
	file, err := a.ginCtx.FormFile(key)
	if err != nil {
		panic(dto_response.NewInternalServerErrorResponse())
	}

	return file
}

func (a *apiContext) getParam(key string) string {
	return a.ginCtx.Param(key)
}

func (a *apiContext) getUuidParam(key string) string {
	uuidParam := a.getParam(key)

	if !util.IsUuid(uuidParam) {
		panic(dto_response.NewBadRequestErrorResponse(fmt.Sprintf("%s must be a valid UUID", strcase.ToCamel(key))))
	}

	return uuidParam
}

func (a *apiContext) getUser() model.User {
	return model.MustGetUserCtx(a.context())
}

func (a *apiContext) getUserId() string {
	return a.getUser().Id
}

func (a *apiContext) shouldBind(obj interface{}) error {
	return util.ShouldGinBind(a.ginCtx, obj)
}

func (a *apiContext) mustBind(obj interface{}) {
	if err := a.shouldBind(obj); err != nil {
		panic(a.translateBindErr(err))
	}
}

func (a *apiContext) translateBindErr(err error) dto_response.ErrorResponse {
	var r dto_response.ErrorResponse

	switch v := err.(type) {
	case validator.StructValidationErrors:
		errs := []dto_response.Error{}
		translations := v.Translate(model.MustGetValidatorTranslatorCtx(a.context()))
		for k, translation := range translations {
			errs = append(errs, dto_response.Error{
				Domain:  k,
				Message: translation,
			})
		}

		r = dto_response.NewBadRequestErrorResponse("Invalid request payload")
		r.Errors = errs

	case *json.UnmarshalTypeError, *json.InvalidUnmarshalError:
		r = dto_response.NewBadRequestErrorResponse("Invalid request payload")

	default:
		switch v {
		case bindingInternal.ErrConvertMapStringSlice, bindingInternal.ErrConvertToMapString,
			bindingInternal.ErrMultiFileHeader, bindingInternal.ErrMultiFileHeaderLenInvalid,
			bindingInternal.ErrIgnoredBinding:
			r = dto_response.NewBadRequestErrorResponse("Invalid request payload")

		default:
			panic(err)
		}
	}

	return r
}

func (a *apiContext) dataFromReader(code int, contentLength int64, contentType string, reader io.Reader, extraHeaders map[string]string) {
	a.ginCtx.DataFromReader(
		code,
		contentLength,
		contentType,
		reader,
		extraHeaders,
	)
}

func (a *apiContext) json(code int, obj interface{}) {
	a.ginCtx.JSON(code, obj)
}

type api struct {
	permissionUseCase use_case.PermissionUseCase
}

func newApi(useCaseManager use_case.UseCaseManager) api {
	return api{
		permissionUseCase: useCaseManager.PermissionUseCase(),
	}
}

func (a *api) Authorize(permissionEnum *data_type.Permission, fn func(ctx apiContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apiCtx := newApiContext(ctx)

		if permissionEnum != nil {
			a.permissionUseCase.Authorize(
				apiCtx.context(),
				*permissionEnum,
				apiCtx.getClientIp(),
			)
		}

		fn(apiCtx)
	}
}

func (a *api) Guest(fn func(ctx apiContext)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(newApiContext(ctx))
	}
}

type customHttpDir http.Dir

func (d customHttpDir) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	if strings.HasPrefix(path.Clean("/"+name), path.Clean("/"+constant.PrivatePath)) {
		return nil, fs.ErrNotExist
	}

	httpDir := http.Dir(d)

	return httpDir.Open(name)
}

func registerMiddlewares(router gin.IRouter, container *manager.Container) {
	useCaseManager := container.UseCaseManager()

	middleware.ApiVersionHandler(router)
	middleware.RequestIdHandler(router)
	middleware.TranslatorHandler(router)
	middleware.PanicHandler(router, container.LoggerStack())
	middleware.JWTHandler(router, useCaseManager.AuthUseCase())
}

func registerRoutes(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	if global.GetFilesystem() == filesystemInternal.FilesystemLocal {
		router.StaticFS(
			fmt.Sprintf("/%s", global.StaticStorageFs),
			customHttpDir(fmt.Sprintf("%s/%s", global.GetStorageDir(), global.StaticStorageFs)),
		)
	}

	RegisterAdminBalanceApi(router, useCaseManager)
	RegisterAdminUserApi(router, useCaseManager)
	RegisterAdminUnitApi(router, useCaseManager)
	RegisterAdminProductUnitApi(router, useCaseManager)

	RegisterCustomerApi(router, useCaseManager)
	RegisterProductReceiveApi(router, useCaseManager)
	RegisterSupplierTypeApi(router, useCaseManager)
	RegisterSupplierApi(router, useCaseManager)

	RegisterAuthApi(router, useCaseManager)
}

func NewRouter(container *manager.Container) *gin.Engine {
	allowedHeaders := []string{
		"Accept",
		"Accept-Encoding",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Content-Length",
		"Origin",
		"X-CSRF-Token",
		"X-Requested-With",
	}

	if global.IsProduction() || global.IsTesting() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins: global.GetConfig().CorsAllowedOrigins,
				AllowMethods: []string{
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodHead,
				},
				AllowHeaders: allowedHeaders,
				ExposeHeaders: []string{
					"Content-Type",
					"Content-Length",
					"Content-Disposition",
				},
				AllowCredentials: true,
				MaxAge:           2 * time.Hour,
			},
		),
	)

	router.ForwardedByClientIP = false
	if global.IsProduction() {
		router.TrustedPlatform = "CloudFront-Viewer-Address"
	} else {
		router.TrustedPlatform = "X-GWS-IP"
	}

	registerMiddlewares(router, container)

	registerRoutes(router, container.UseCaseManager())

	return router
}
