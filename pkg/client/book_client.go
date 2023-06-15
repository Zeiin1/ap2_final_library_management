package client

import (
	"context"
	"fmt"
	"github.com/sat0urn/go-grpc-order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type BookServiceClient struct {
	Client pb.BookServiceClient
}

func InitBookServiceClient(url string) BookServiceClient {
	cc, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := BookServiceClient{
		Client: pb.NewBookServiceClient(cc),
	}

	return c
}

func (c *BookServiceClient) FindOne(productId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: productId,
	}

	return c.Client.FindOne(context.Background(), req)
}

func (c *BookServiceClient) DecreaseStock(productId int64, orderId int64) (*pb.DecreaseStockResponse, error) {
	req := &pb.DecreaseStockRequest{
		Id:      productId,
		OrderId: orderId,
	}

	return c.Client.DecreaseStock(context.Background(), req)
}
