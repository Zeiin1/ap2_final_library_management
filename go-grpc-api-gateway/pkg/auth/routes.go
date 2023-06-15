package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}
	r.Static("/static", "C:\\Users\\User\\GolandProjects\\ap2_final_library_management\\go-grpc-api-gateway\\pkg\\web\\static")
	r.GET("/", svc.ShowMainPage)
	r.GET("/register", svc.ShowRegisterPage)
	r.GET("/loginPage", svc.ShowLoginPage)
	r.POST("/register", svc.Register)
	r.POST("/login", svc.Login)

	return svc
}
func (svc *ServiceClient) ShowLoginPage(ctx *gin.Context) {
	w := ctx.Writer
	routes.ShowLoginPage(w, ctx, svc.Client)
}
func (svc *ServiceClient) ShowMainPage(ctx *gin.Context) {
	w := ctx.Writer
	routes.ShowMainPage(w, ctx, svc.Client)
}
func (svc *ServiceClient) ShowRegisterPage(ctx *gin.Context) {
	w := ctx.Writer
	routes.ShowRegisterPage(w, ctx, svc.Client)
}
func (svc *ServiceClient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceClient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}
