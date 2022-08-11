package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/auth"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/config"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/front"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/order"
	"github.com/heifetzwu/go-grpc-api-gateway/pkg/product"
)

func main() {
	fmt.Println("###0 main 1")
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)
	front.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
