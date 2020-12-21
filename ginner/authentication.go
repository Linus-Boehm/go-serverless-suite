package ginner

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Linus-Boehm/go-serverless-suite/common"

	"github.com/gin-gonic/gin"
)

const UserClaimsContext = "User"

type AuthenticationServicer interface {
	GetClaimsFromToken(accessToken string) (User, error)
}

type RoleReceiver interface {
	String() string
}

type User interface {
	HasGroup(roles RoleReceiver) bool
}

func AuthenticationMiddleware(service AuthenticationServicer) gin.HandlerFunc {
	return func(context *gin.Context) {
		var user interface{}
		authHeader := context.GetHeader("Authorization")
		const BearerSchema = "Bearer"
		if len(authHeader) > len(BearerSchema) {
			tokenString := strings.TrimSpace(authHeader[len(BearerSchema):])
			claims, err := service.GetClaimsFromToken(tokenString)
			if err != nil {
				context.Error(common.NewNotAuthorizedError())
				return
			}
			user = claims
		}
		context.Set(UserClaimsContext, user)
	}
}

func RequireAuthenticationMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, err := GetUserFromContext(context)
		if err != nil || user == nil {
			context.Error(errors.Wrap(common.NewNotAuthorizedError(), fmt.Sprint(err)))
		}
	}
}

func RequireRolesMiddleware(roles ...RoleReceiver) gin.HandlerFunc {
	return func(context *gin.Context) {
		user, err := GetUserFromContext(context)
		if err != nil || user == nil {
			context.Error(errors.Wrap(common.NewNotAuthorizedError(), fmt.Sprint(err)))
			return
		}
		var rolesStr string
		for _, role := range roles {
			if user.HasGroup(role) {
				return

			}
			if rolesStr == "" {
				rolesStr = role.String()
			} else {
				rolesStr += fmt.Sprint(", ", role.String())
			}

		}
		err = common.NewInsufficientPermissionsError(fmt.Sprintf("User needs at least one of role: %s", rolesStr))
		context.Error(err)
	}
}

func GetUserFromContext(ctx *gin.Context) (User, error) {
	rawUser, exists := ctx.Get(UserClaimsContext)
	if !exists {
		return nil, nil
	}
	var userClaims User
	userClaims, ok := rawUser.(User)
	if !ok {
		return nil, errors.New("Corrupted User information")
	}
	return userClaims, nil
}
