package product

import (
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/auth"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/config"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	rts := r.Group("/product")
	rts.Use(a.AuthRequired)
	rts.POST("/", svc.CreateProduct)
	rts.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
