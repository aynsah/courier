package controllers

import (
	"courier/pkg/response"
	"courier/services/logistic/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogisticList(c *gin.Context) {
	var logistics []models.Logistic
	status, err := models.GetAllLogistics(&logistics)
	if err != nil {
		response.ShowErr(c, status, "Can't find data logistic", err.Error())
		return
	}

	c.JSON(status, gin.H{
		"status":    status,
		"logistics": logistics,
	})
}

func SearchLogistics(c *gin.Context) {
	origin_name := c.PostForm("origin_name")
	destination_name := c.PostForm("destination_name")

	if origin_name == "" || destination_name == "" {
		response.ShowErr(c, http.StatusBadRequest, "Value can't be nil", nil)
	}

	var logistics []models.Logistic
	status, err := models.FindLogistics(origin_name, destination_name, &logistics)
	if err != nil {
		response.ShowErr(c, status, "Can't find data logistic", err.Error())
		return
	}

	c.JSON(status, gin.H{
		"status":    status,
		"logistics": logistics,
	})
}
