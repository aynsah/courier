package main

import (
	"courier/services/user/config"
	"courier/services/user/controllers"
	"courier/services/user/database"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

const validToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtc2lzZG4iOiI2Mjg5MTYyMzUxMjMiLCJ1c2VybmFtZSI6Imd1bWl5YSIsInV1aWQiOiIwMmM0Yzc1ZC0zNDRlLTExZWMtOGYzNi0wMDE1NWQwNzRjOWUifQ.0uhFfOYEBxko1iVwxieo49F0T5gDzgleJrj-H__KyDU"
const invalidToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IpXVCJ9.eyJtc2lzZG4iOiI2Mjg5MTYyMzUxMjMiLCJ1c2VybmFtZSI6Imd1bWl5YSIsInV1aWQiOiIwMmM0Yzc1ZC0zNDRlLTExZWMtOGYzNi0wMDE1NWQwNzRjOWUifQ.0uhFfOYEBxko1iVwxieo49F0T5gDzgleJrj-H__KyDU"

func initServer() {
	config.LoadConfig()
	database.LoadDatabase()
}

func TestLogin(t *testing.T) {
	initServer()
	parm := url.Values{}
	parm.Add("MSISDN", "6281221501224")
	parm.Add("Password", "wellsad")

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/v1/users/login", controllers.Login)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/login", strings.NewReader(parm.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("Error while Login: %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusUnauthorized {
		t.Logf("Expected result: Wrong MSISIDN/Password")
	} else if w.Code != http.StatusOK {
		t.Fatalf("Error code while Login: %v\n", w.Code)
	}

	t.Logf("Expected result: login success")
}

func TestInfoValidToken(t *testing.T) {
	initServer()

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/v1/users/info", controllers.Info)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/info", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", validToken)
	if err != nil {
		t.Fatalf("Error while getting info : %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Error code getting info: %v\n", w.Code)
	}

	t.Logf("Expected result: getting info success")
}

func TestInfoInvalidToken(t *testing.T) {
	initServer()

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/v1/users/info", controllers.Info)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/info", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", invalidToken)
	if err != nil {
		t.Fatalf("Error while getting info : %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Status code getting info: %v\n", w.Code)
	}

	t.Logf("Expected result: token invalid")
}

func TestSignUpValid(t *testing.T) {
	initServer()
	parm := url.Values{}
	parm.Add("MSISDN", "6281222341234")
	parm.Add("Username", "NewUsername")
	parm.Add("Password", "NewPassword")

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/v1/users/sign-up", controllers.SignUp)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/sign-up", strings.NewReader(parm.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("Error while sign up : %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code == http.StatusForbidden {
		t.Logf("Expected result: sign up failed (Username or MSISDN has been used)")
	} else if w.Code != http.StatusOK {
		t.Fatalf("Error code sign up: %v\n", w.Body)
	}

	t.Logf("Expected result: sign up success")
}

func TestSignUpInvalid(t *testing.T) {
	initServer()
	parm := url.Values{}
	parm.Add("MSISDN", "081222341234")
	parm.Add("Username", "NewUsername2")
	parm.Add("Password", "NewPassword2")

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.POST("/api/v1/users/sign-up", controllers.SignUp)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/users/sign-up", strings.NewReader(parm.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatalf("Error while sign up : %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("Status code sign up: %v\n", w.Code)
	}

	t.Logf("Expected result: sign up invalid (MSISDN must start with 62)")
}
