package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/common"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	fmt.Println("###1  auth required")
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
	fmt.Println("### 3: token = ", token[1])
	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	fmt.Println("### 4: res = ", res)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}

func (c *AuthMiddlewareConfig) AuthRequiredToken(ctx *gin.Context) {
	fmt.Println("###1  auth required token")

	token := common.ReadToken(ctx)

	if len(token) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println("### token = ", token)

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token,
	})

	fmt.Println("### 4: res = ", res)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", res.UserId)

	ctx.Next()
}

func CORS(ctx *gin.Context) {
	fmt.Println("### CORS ctx.Request.Method = ", ctx.Request.Method)
	// ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:9090")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.Writer.Header().Set("Access-Control-Allow-Headers", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if ctx.Request.Method == "OPTIONS" {

		fmt.Println("### method = options")
		ctx.AbortWithStatus(http.StatusNoContent)
	}
	ctx.Next()

}
