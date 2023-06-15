package routes

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sat0urn/go-grpc-api-gateway/pkg/auth/pb"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type User struct {
	Name     string `json:"name"`
	SurName  string `json:"sur_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

/*
	func (c *gin.Context) ParseForm() error {
		if c.Request.Form == nil {
			// Parse the form data
			err := c.Request.ParseForm()
			if err != nil {
				return err
			}

			// Set the parsed form data to the Gin context
			c.Request.PostForm = c.Request.Form
			c.Request.Form = nil
		}

		return nil
	}
*/
func Register(ctx *gin.Context, c pb.AuthServiceClient) {

	name := ctx.PostForm("name")
	surname := ctx.PostForm("surname")
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	body := User{Name: name,
		SurName:  surname,
		Email:    email,
		Password: password}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Name:     body.Name,
		SurName:  body.SurName,
		Email:    body.Email,
		Password: body.Password,
	})

	location := url.URL{Path: "/loginPage"}
	ctx.Redirect(http.StatusFound, location.RequestURI())

	log.Println("=======user: ", body.Name)

	if err != nil {
		log.Println(res.Error)
		return
	}

	ctx.JSON(200, &res)
}
func ShowMainPage(w http.ResponseWriter, ctx *gin.Context, c pb.AuthServiceClient) {
	ts, err := template.ParseFiles("C:\\Users\\User\\GolandProjects\\ap2_final_library_management\\go-grpc-api-gateway\\pkg\\web\\templates\\index.html")

	if err != nil {
		log.Println(err.Error())

	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())

	}
}

func ShowRegisterPage(w http.ResponseWriter, ctx *gin.Context, c pb.AuthServiceClient) {
	user := User{}
	ts, err := template.ParseFiles("C:\\Users\\User\\GolandProjects\\ap2_final_library_management\\go-grpc-api-gateway\\pkg\\web\\templates\\reg.html")

	if err != nil {
		log.Println(err.Error())

		return
	}
	err = ts.Execute(w, user)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
