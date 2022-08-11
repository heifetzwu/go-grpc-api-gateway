package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	key   = []byte("1234")
	store = sessions.NewCookieStore(key)
	// store = sessions.NewCookieStore()
	// store = sessions.NewCookie("add", "", nil)
	// store = sessions.CookieStore(sessions)
)

func WriteToken(ctx *gin.Context, t string) {
	fmt.Println("WriteToken")
	store.Options = &sessions.Options{
		MaxAge:   60 * 5,
		HttpOnly: false,
		Path:     "/",
	}
	session, err := store.Get(ctx.Request, "token")
	if err != nil {
		fmt.Println("store.Get err")
	}

	session.Values["tokenv"] = t
	session.Values["token1"] = 1
	session.Save(ctx.Request, ctx.Writer)
}

func ReadToken(ctx *gin.Context) string {
	session, _ := store.Get(ctx.Request, "token")
	t := session.Values["tokenv"]
	t1 := session.Values["token1"]
	fmt.Println("### t1 = ", t1)
	var token string
	if t == nil {
		token = ""
	} else {
		token = session.Values["tokenv"].(string)
	}
	return token
}
