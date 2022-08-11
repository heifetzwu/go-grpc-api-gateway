package front

import (
	"github.com/gin-gonic/gin"

	"github.com/heifetzwu/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) {
	// r.LoadHTMLGlob("/home/jack/user/workspaces/go-grpc-project/go-grpc-api-gateway/pkg/front/html/*")
	// r.LoadHTMLGlob("html/*")
	r.LoadHTMLGlob("pkg/front/html/*")
	routes := r.Group("/front")
	routes.GET("/login2", LoginGet)

}
