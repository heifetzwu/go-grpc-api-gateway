package product

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/auth"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/config"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) {
	fmt.Print("###2 RegisterRoutes authRequired")
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := r.Group("/product")
	routes.Use(auth.CORS)
	// routes.Use(a.AuthRequired)
	routes.Use(a.AuthRequiredToken)

	routes.POST("/", svc.CreateProduct)
	routes.GET("/:id", svc.FindOne)
	routes.OPTIONS("/:id")
}

func (svc *ServiceClient) FindOne(ctx *gin.Context) {
	fmt.Println("### ctx.Request.URL=", ctx.Request.URL)
	fmt.Println("### ctx.Request.URL.Host=", ctx.Request.URL.Host)
	fmt.Println("### ctx.Request.Host=", ctx.Request.Host)
	fmt.Println("### ctx.Request=", ctx.Request)
	routes.FineOne(ctx, svc.Client)
}

func (svc *ServiceClient) CreateProduct(ctx *gin.Context) {
	routes.CreateProduct(ctx, svc.Client)
}
