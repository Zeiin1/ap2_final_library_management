package main

import (
	"fmt"
	"github.com/sat0urn/go-grpc-order-svc/pkg/client"
	"github.com/sat0urn/go-grpc-order-svc/pkg/config"
	"github.com/sat0urn/go-grpc-order-svc/pkg/db"
	"github.com/sat0urn/go-grpc-order-svc/pkg/pb"
	service "github.com/sat0urn/go-grpc-order-svc/pkg/services"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	bookSvc := client.InitBookServiceClient(c.BookSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := service.Server{
		H:       h,
		BookSvc: bookSvc,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to server:", err)
	}
}
