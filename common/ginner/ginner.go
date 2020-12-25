package ginner

import (
	"net/http"

	"github.com/Linus-Boehm/go-serverless-suite/common"
	"github.com/gin-contrib/logger"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	"github.com/gin-gonic/gin"
)

type ErrorType string

// List of ErrorType
const (
	EntityContext                 string    = "ENTITY"
	ErrorTypeEntityNotFound       ErrorType = "EntityNotFound"
	ErrorTypeEntityInvalid        ErrorType = "EntityInvalid "
	ErrorTypeEntityAlreadyExists  ErrorType = "EntityAlreadyExists"
	ErrorTypeEntityNotEmpty       ErrorType = "EntityNotEmpty"
	ErrorTypeAuthentication       ErrorType = "Authentication"
	ErrorTypeMethodNotImplemented ErrorType = "ErrorTypeMethodNotImplemented"
	ErrorTypeInternal             ErrorType = "Internal"
)

type GenericError struct {
	Entity     string    `json:"entity"`
	StatusCode int       `json:"statusCode"`
	ErrorCode  ErrorType `json:"errorCode"`
	Message    string    `json:"message"`
}

type GinErrorMapper func(err error, ctx *gin.Context) *GenericError

func ErrInsufficientPermissionMapper(err error, ctx *gin.Context) *GenericError {
	e := common.NewInsufficientPermissionsError("")
	if errors.Is(err, e) {
		return &GenericError{
			Entity:     getEntityFromCtx(ctx),
			StatusCode: http.StatusForbidden,
			ErrorCode:  ErrorTypeAuthentication,
			Message:    e.Error(),
		}
	}
	return nil
}

func ErrNotAuthorizedMapper(err error, ctx *gin.Context) *GenericError {
	e := common.NewNotAuthorizedError()
	if errors.Is(err, e) {
		return &GenericError{
			Entity:     getEntityFromCtx(ctx),
			StatusCode: http.StatusUnauthorized,
			ErrorCode:  ErrorTypeAuthentication,
			Message:    e.Error(),
		}
	}
	return nil
}

func NewGiner(cfg common.Configer, zerologger *zerolog.Logger, errorMapper ...GinErrorMapper) *gin.Engine {
	if cfg.GetStage() == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.New()

	// Custom logger
	router.Use(logger.SetLogger(logger.Config{
		Logger: zerologger,
		UTC:    true,
	}))
	router.Use(CorsMiddleware())
	// this should be added last, as it watches next
	router.Use(errorMapperHandler(errorMapper...))
	return router
}

func errorMapperHandler(mapper ...GinErrorMapper) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		detectedErrs := ctx.Errors.ByType(gin.ErrorTypeAny)
		if len(detectedErrs) == 0 {
			return
		}
		// first error in chain
		err := detectedErrs[0].Err
		for _, mapErr := range mapper {
			gErr := mapErr(err, ctx)
			if gErr != nil {
				ctx.JSON(gErr.StatusCode, gErr)
				return
			}
		}
		defaultErr := &GenericError{
			Entity:     getEntityFromCtx(ctx),
			StatusCode: http.StatusInternalServerError,
			ErrorCode:  ErrorTypeInternal,
			Message:    "Internal Server Error",
		}
		ctx.JSON(defaultErr.StatusCode, defaultErr)
	}
}

func getEntityFromCtx(ctx *gin.Context) string {
	if val, exist := ctx.Get(EntityContext); exist {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return "UNKNOWN"
}
