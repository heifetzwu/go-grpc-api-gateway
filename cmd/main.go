package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/auth"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/config"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/order"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/product"
)

func main() {
	fmt.Println("###0 main")
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
