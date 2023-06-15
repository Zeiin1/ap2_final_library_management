package services

import (
	"context"
	"github.com/sat0urn/go-grpc-auth-svc/pkg/db"
	"github.com/sat0urn/go-grpc-auth-svc/pkg/models"
	"github.com/sat0urn/go-grpc-auth-svc/pkg/pb"
	"github.com/sat0urn/go-grpc-auth-svc/pkg/utils"
	"log"
)

type Server struct {
	H   db.Handler
	Jwt utils.JwtWrapper
}

func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User
	log.Println("==================")
	result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user)

	if result != nil {
		log.Println(result.Error)
	}

	user.Password = utils.HashPassword(req.Password)
	user.SurName = req.SurName
	user.Name = req.Name
	user.Email = req.Email
	log.Println("surname= ", req.SurName)

	s.H.DB.Create(&user)
	return nil, nil

}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user)
	if result != nil {
		log.Println(result.Error)
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)

	if !match {
		return &pb.LoginResponse{Error: "some",
			Id: 0}, nil
	}

	id := user.Id

	return &pb.LoginResponse{Error: "",
		Id: int32(id)}, nil
}

/*func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {


	var user models.User

	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil
}*/
