package main

import (
	"courier/services/logistic/config"
	"courier/services/logistic/controllers"
	"courier/services/logistic/database"
	"courier/services/logistic/middleware"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func initServer() {
	config.LoadConfig()
	database.LoadDatabase()
}

func TestLogisticListValidToken(t *testing.T) {
	initServer()

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(middleware.Verify)
	r.POST("/api/v1/logistics/", controllers.LogisticList)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/logistics/", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtc2lzZG4iOiI2Mjg5MTYyMzUxMjMiLCJ1c2VybmFtZSI6Imd1bWl5YSIsInV1aWQiOiIwMmM0Yzc1ZC0zNDRlLTExZWMtOGYzNi0wMDE1NWQwNzRjOWUifQ.0uhFfOYEBxko1iVwxieo49F0T5gDzgleJrj-H__KyDU")
	if err != nil {
		t.Fatalf("Error while getting list : %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Error code getting list: %v\n", w.Code)
	}

	t.Logf("Expected result: getting list success")
}

func TestLogisticListInvalidToken(t *testing.T) {
	initServer()

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(middleware.Verify)
	r.POST("/api/v1/logistics/", controllers.LogisticList)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/logistics/", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IpXVCJ9.eyJtc2lzZG4iOiI2Mjg5MTYyMzUxMjMiLCJ1c2VybmFtZSI6Imd1bWl5YSIsInV1aWQiOiIwMmM0Yzc1ZC0zNDRlLTExZWMtOGYzNi0wMDE1NWQwNzRjOWUifQ.0uhFfOYEBxko1iVwxieo49F0T5gDzgleJrj-H__KyDU")
	if err != nil {
		t.Fatalf("Error while getting list : %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Fatalf("Status code getting list: %v\n", w.Code)
	}

	t.Logf("Expected result: Invalid token while getting list")
}

func TestSearchLogistics(t *testing.T) {
	initServer()
	parm := url.Values{}
	parm.Add("origin_name", "sad")
	parm.Add("destination_name", "sda")

	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.Use(middleware.Verify)
	r.POST("/api/v1/logistics/search", controllers.SearchLogistics)

	req, err := http.NewRequest(http.MethodPost, "/api/v1/logistics/search", strings.NewReader(parm.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJtc2lzZG4iOiI2Mjg5MTYyMzUxMjMiLCJ1c2VybmFtZSI6Imd1bWl5YSIsInV1aWQiOiIwMmM0Yzc1ZC0zNDRlLTExZWMtOGYzNi0wMDE1NWQwNzRjOWUifQ.0uhFfOYEBxko1iVwxieo49F0T5gDzgleJrj-H__KyDU")
	if err != nil {
		t.Fatalf("Error while getting list by origin name and destination name: %v\n", err.Error())
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Error code getting list by origin name and destination name: %v\n", w.Code)
	}

	t.Logf("Expected result: getting list by origin name and destination name success")
}
