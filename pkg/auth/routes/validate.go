package routes

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/auth/pb"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/common"
)

func ValidateCookie(ctx *gin.Context, c pb.AuthServiceClient) {

	fmt.Println("###1  auth required token")

	token := common.ReadToken(ctx)

	if len(token) == 0 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println("### token = ", token)

	res, err := c.Validate(context.Background(), &pb.ValidateRequest{
		Token: token,
	})

	fmt.Println("### 4: res = ", res)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Println("### res.GetStatus()=", res.GetStatus(), res.GetError())
	if res.GetStatus() > 300 {
		ctx.AbortWithError(int(res.GetStatus()), errors.New(res.GetError()))

		return
	}
	// common.WriteToken(ctx, res.GetToken())
	ctx.JSON(http.StatusAccepted, &res)
}
