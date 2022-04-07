package middleware

import (
	"edx_go/common"
	app_ctx "edx_go/component"
	tokenprovider "edx_go/providers/token"
	"errors"
	"github.com/gin-gonic/gin"
	"strings"
)

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("Authorization not valid")
	}

	return parts[1], nil
}

func AuthMiddleware(appCtx app_ctx.AppContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

		if err != nil {
			panic(common.ErrNoPermission(err))
			return
		}

		userId, err := tokenprovider.ValidateToken(token)

		if err != nil {
			panic(common.ErrNoPermission(err))
			return
		}

		ctx.Set("user_id", userId)
		ctx.Next()

	}
}
