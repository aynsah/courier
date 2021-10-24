package controllers

import (
	"courier/pkg/response"
	"courier/services/logistic/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogisticList godoc
// @Summary List of Logistics
// @Schemes
// @Description Get all Logistics from database
// @Tags Logistic
// @Accept multipart/form-data
// @Produce json
// @Success 200 {array} []models.Logistic
// @Failder 401 {string} message
// @Failder 500 {string} message
// @Router / [post]
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

// SearchLogistics godoc
// @Summary Search Data Logistics
// @Schemes
// @Description Search Logistics from database by origin_name and destination_name
// @Tags Logistic
// @Accept multipart/form-data
// @Produce json
// @Success 200 {array} []models.Logistic
// @Failder 400 {string} message
// @Failder 401 {string} message
// @Failder 500 {string} message
// @Router /search [post]
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
