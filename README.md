```go
package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

//database is working in docker
func TestShowIndexPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Errorf("could not get the healthcheck endpoint")
	}
	r := httptest.NewRecorder()
	handler := gin.Engine{}
	handler.ServeHTTP(r, req)
	status := r.Code
	urlString := r.Header().Get("Location")
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "", urlString)
}

func TestShowRegistrationPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/register", nil)
	if err != nil {
		t.Errorf("could not get the healthcheck endpoint")
	}
	r := httptest.NewRecorder()
	handler := gin.Engine{}
	handler.ServeHTTP(r, req)
	status := r.Code
	urlString := r.Header().Get("Location")
	assert.Equal(t, http.StatusOK, status)
	assert.Equal(t, "", urlString)
}

func TestServiceClient_ShowLoginPage(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Errorf("could not get the healthcheck endpoint")
	}
	r := httptest.NewRecorder()
	handler := gin.Engine{}
	handler.ServeHTTP(r, req)
	status := r.Code
	urlString := r.Header().Get("Location")
	assert.Equal(t, http.StatusSeeOther, status)
	assert.Equal(t, "", urlString)
}

func TestRegisterUserHandler(t *testing.T) {
	req, err := http.NewRequest("POST", "/register?name=Zeiin&surname=Krasava&email=zeyn03@gmail.com&password=123456", nil)
	if err != nil {
		t.Errorf("could not get the registration endpoint")
	}
	r := httptest.NewRecorder()
	handler := gin.Engine{}
	handler.ServeHTTP(r, req)
	status := r.Code
	urlString := r.Header().Get("Location")
	assert.Equal(t, http.StatusSeeOther, status)
	assert.Equal(t, "/loginPage", urlString)
}

func TestUserLogin(t *testing.T) {
	req, err := http.NewRequest("POST", "/login?email=zeyn03@gmail.com&password=123456", nil)
	if err != nil {
		t.Errorf("could not get the registration endpoint")
	}
	r := httptest.NewRecorder()
	handler := gin.Engine{}
	handler.ServeHTTP(r, req)
	status := r.Code
	urlString := r.Header().Get("Location")
	assert.Equal(t, http.StatusSeeOther, status)

	assert.Equal(t, "/profile/28", urlString)
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/delete/28", nil)
	if err != nil {
		log.Print(err)
	}

	url := req.URL.String()

	r := httptest.NewRecorder()

	handler := gin.Engine{}
	handler.ServeHTTP(r, req)
	println("url is:", url)
	status := r.Code
	urlString := r.Header().Get("Location")
	assert.Equal(t, http.StatusSeeOther, status)
	assert.Equal(t, "/", urlString)
}

func TestUpdateUserInfo(t *testing.T) {
	req, err := http.NewRequest("POST", `/update/28?name=updatedName&surname=Krasava&email=zeyn03@gmail.com&password=123456`, nil)

	if err != nil {
		t.Errorf("could not get the registration endpoint")
	}
	r := httptest.NewRecorder()

	handler := gin.Engine{}
	handler.ServeHTTP(r, req)
	status := r.Code
	urlString := r.Header().Get("Location")
	assert.Equal(t, http.StatusSeeOther, status)
	assert.Equal(t, "/user_info/28", urlString)
}
```
