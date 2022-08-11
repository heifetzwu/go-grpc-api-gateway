package routes

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/product/pb"
)

var (
	key   = []byte("1234")
	store = sessions.NewCookieStore(key)
	// store = sessions.NewCookieStore()
	// store = sessions.NewCookie("add", "", nil)
	// store = sessions.CookieStore(sessions)
)

func FineOne(ctx *gin.Context, c pb.ProductServiceClient) {
	v, _ := ctx.Get("userId")
	fmt.Println("### FindOne", v)
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 32)

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	addcookie(ctx)
	ctx.JSON(http.StatusCreated, &res)
}

func addcookie(ctx *gin.Context) {
	fmt.Println("addcookie")
	store.Options = &sessions.Options{
		MaxAge:   60 * 5,
		HttpOnly: false,
		Path:     "/",
	}
	session, err := store.Get(ctx.Request, "addcookie")
	if err != nil {
		fmt.Println("store.Get err")
	}
	session.Values["ckbook"] = true
	session.Values["ckint"] = 1
	session.Values["ckstr"] = "today20220507"
	session.Save(ctx.Request, ctx.Writer)

}
