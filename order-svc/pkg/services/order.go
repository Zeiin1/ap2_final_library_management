package service

import (
	"context"
	"github.com/sat0urn/go-grpc-order-svc/pkg/client"
	"github.com/sat0urn/go-grpc-order-svc/pkg/db"
	"github.com/sat0urn/go-grpc-order-svc/pkg/models"
	"github.com/sat0urn/go-grpc-order-svc/pkg/pb"
	"net/http"
)

type Server struct {
	H       db.Handler
	BookSvc client.BookServiceClient
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	book, err := s.BookSvc.FindOne(req.BookId)

	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if book.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{
			Status: book.Status,
			Error:  book.Error,
		}, nil
	} else if book.Data.Stock < req.Quantity {
		return &pb.CreateOrderResponse{
			Status: http.StatusConflict,
			Error:  "Stock too less",
		}, nil
	}

	order := models.Order{
		Price:  book.Data.Price,
		BookId: book.Data.Id,
		UserId: req.UserId,
	}

	s.H.DB.Create(&order)

	res, err := s.BookSvc.DecreaseStock(req.BookId, order.Id)
	if err != nil {
		return &pb.CreateOrderResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{
			Status: http.StatusContinue,
			Error:  res.Error,
		}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id:     order.Id,
	}, nil
}
