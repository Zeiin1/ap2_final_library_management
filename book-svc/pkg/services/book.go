package services

import (
	"context"
	"github.com/sat0urn/book-svc/pkg/db"
	"github.com/sat0urn/book-svc/pkg/models"
	"github.com/sat0urn/book-svc/pkg/pb"
	"net/http"
)

type Server struct {
	H db.Handler
}

func (s *Server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	var book models.Book

	book.Name = req.Name
	book.Stock = req.Stock
	book.Price = req.Price

	if result := s.H.DB.Create(&book); result.Error != nil {
		return &pb.CreateBookResponse{
			Status: http.StatusConflict,
			Error:  result.Error.Error(),
		}, nil
	}

	return &pb.CreateBookResponse{
		Status: http.StatusCreated,
		Id:     book.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var book models.Book

	if result := s.H.DB.First(&book, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	data := &pb.FindOneData{
		Id:    book.Id,
		Name:  book.Name,
		Stock: book.Stock,
		Price: book.Stock,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   data,
	}, nil
}

func (s *Server) DecreaseStock(ctx context.Context, req *pb.DecreaseStockRequest) (*pb.DecreaseStockResponse, error) {
	var book models.Book

	if result := s.H.DB.First(&book, req.Id); result.Error != nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusNotFound,
			Error:  result.Error.Error(),
		}, nil
	}

	if book.Stock <= 0 {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock too low",
		}, nil
	}

	var log models.StockDecreaseLog

	if result := s.H.DB.Where(&models.StockDecreaseLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DecreaseStockResponse{
			Status: http.StatusConflict,
			Error:  "Stock already decreased",
		}, nil
	}

	book.Stock = book.Stock - 1

	s.H.DB.Save(&book)

	log.OrderId = req.OrderId
	log.BookRefer = book.Id

	s.H.DB.Create(&log)

	return &pb.DecreaseStockResponse{
		Status: http.StatusOK,
	}, nil
}
