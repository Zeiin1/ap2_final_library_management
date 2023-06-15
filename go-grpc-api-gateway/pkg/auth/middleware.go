package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{
		svc: svc,
	}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Next()
}
