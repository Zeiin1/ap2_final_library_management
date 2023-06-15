package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/auth/pb"
	"html/template"
	"log"
	"net/http"
)

type LoginRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {

	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    email,
		Password: password,
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}
	if res.Error != "" {
		ctx.Redirect(http.StatusSeeOther, "/")

	}

	ctx.JSON(http.StatusCreated, &res)
}
func ShowLoginPage(w http.ResponseWriter, ctx *gin.Context, c pb.AuthServiceClient) {

	ts, err := template.ParseFiles("C:\\Users\\User\\GolandProjects\\ap2_final_library_management\\go-grpc-api-gateway\\pkg\\web\\templates\\login.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
