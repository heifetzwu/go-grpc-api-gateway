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

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    b.Email,
		Password: b.Password,
	})

	fmt.Println("#### login ", b.Email, b.Password, err)
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	fmt.Println("### res.GetStatus()=", res.GetStatus(), res.GetError())
	if res.GetStatus() > 300 {
		ctx.AbortWithError(int(res.GetStatus()), errors.New(res.GetError()))

		return
	}
	common.WriteToken(ctx, res.GetToken())
	ctx.JSON(http.StatusCreated, &res)
}
