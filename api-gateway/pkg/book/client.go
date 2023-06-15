package book

import (
	"fmt"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/book/pb"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.BookServiceClient
}

func InitServiceClient(c *config.Config) pb.BookServiceClient {
	cc, err := grpc.Dial(c.BookSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect: ", err)
	}

	return pb.NewBookServiceClient(cc)
}
