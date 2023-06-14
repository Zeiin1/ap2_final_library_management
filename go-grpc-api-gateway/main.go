package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/auth"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/config"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/order"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/product"
	"log"
)

func main() {
	c, err := config.LoadConfig()
	log.Println(c)
	if err != nil {
		log.Fatalln("Failed at config: ", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	err = r.Run(c.Port)
	if err != nil {
		log.Fatalln("Failed at running: ", err)
	}
}
