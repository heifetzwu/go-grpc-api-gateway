package front

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type LoginRequestBody struct {
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

func LoginGet(ctx *gin.Context) {
	fmt.Println("### login2 ##1")
	fmt.Println("ctx =", ctx)

	ctx.HTML(http.StatusOK, "login.gohtml", gin.H{
		"title": "hello~ jack",
	})

}

func LoginPost(ctx *gin.Context) {
	fmt.Println("### login2 ##1")
	fmt.Println("ctx =", ctx)

	ctx.HTML(http.StatusOK, "login.gohtml", gin.H{
		"title": "hello~ jack",
	})

}

// func Login2(ctx *gin.Context, c pb.AuthServiceClient) {
// 	b := LoginRequestBody{}

// 	if err := ctx.BindJSON(&b); err != nil {
// 		ctx.AbortWithError(http.StatusBadRequest, err)
// 		return
// 	}

// 	res, err := c.Login(context.Background(), &pb.LoginRequest{
// 		Email:    b.Email,
// 		Password: b.Password,
// 	})

// 	if err != nil {
// 		ctx.AbortWithError(http.StatusBadGateway, err)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, &res)
// }
