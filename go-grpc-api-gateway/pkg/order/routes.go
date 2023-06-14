package order

import (
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/auth"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/config"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	rts := r.Group("/order")
	rts.Use(a.AuthRequired)
	rts.POST("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
