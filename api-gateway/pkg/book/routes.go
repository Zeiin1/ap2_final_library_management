package book

import (
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/auth"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/book/routes"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	rts := r.Group("/book")
	rts.Use(a.AuthRequired)
	rts.POST("/", svc.CreateBook)
	rts.GET("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateBook(ctx *gin.Context) {
	routes.CreateBook(ctx, svc.Client)
}
