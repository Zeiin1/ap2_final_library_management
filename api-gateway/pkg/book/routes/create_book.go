package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/book/pb"
	"net/http"
)

type CreateBookRequestBody struct {
	Name  string `json:"name"`
	Stock int64  `json:"stock"`
	Price int64  `json:"price"`
}

func CreateBook(ctx *gin.Context, c pb.BookServiceClient) {
	body := CreateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateBook(context.Background(), &pb.CreateBookRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
