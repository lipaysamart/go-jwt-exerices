package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lipaysamart/go-jwt-exerices/pkg/jtoken"
)

func JWTAuth() gin.HandlerFunc {
	return JWT("x-access")
}

func JWTRefresh() gin.HandlerFunc {
	return JWT("x-refresh")
}

func JWT(tokenType string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		// token := ctx.GetString("Authorization")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}

		payload, err := jtoken.ValidateToken(token)
		if err != nil || payload["type"] != tokenType {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
			return
		}

		ctx.Set("userId", payload["id"])
		ctx.Next()
	}
}
