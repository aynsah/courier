package controllers

import (
	"net/http"

	"courier/pkg/response"
	"courier/pkg/token"
	"courier/services/user/models"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		response.ShowErr(c, http.StatusBadRequest, "Can't bind struct", err.Error())
		return
	}

	if user.MSISDN == "" || user.Username == "" || user.Password == "" {
		response.ShowErr(c, http.StatusBadRequest, "Value cannot be nil", nil)
		return
	}

	if user.MSISDN[:2] != "62" {
		response.ShowErr(c, http.StatusBadRequest, "Invalid MSISDN", nil)
		return
	}

	exists, err := models.IsUserExists(user.MSISDN, user.Username)
	if err != nil {
		response.ShowErr(c, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if exists {
		response.ShowErr(c, http.StatusBadRequest, "User already exists", nil)
		return
	}

	status, err := user.Save()
	if err != nil {
		response.ShowErr(c, status, "Sign up failed", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Sign up success",
	})
}

func Login(c *gin.Context) {
	var credential models.Credential
	err := c.Bind(&credential)
	if err != nil {
		response.ShowErr(c, http.StatusBadRequest, "Can't bind struct", err.Error())
		return
	}

	var user models.User
	status, err := models.FindUserByMSISDN(credential.MSISDN, &user)

	if status == http.StatusUnauthorized {
		response.ShowErr(c, status, "Wrong MSISDN", nil)
		return
	} else if status != http.StatusOK {
		response.ShowErr(c, status, "Login failed", err.Error())
		return
	}

	if user.CheckPassword(credential.Password) != nil {
		response.ShowErr(c, http.StatusUnauthorized, "Wrong Password", nil)
		return
	}

	var JwtToken token.JwtToken
	err = user.GenerateToken(&JwtToken)
	if err != nil {
		response.ShowErr(c, http.StatusInternalServerError, "Error while generating token", err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"token":  JwtToken.Token,
	})
}

func Info(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")

	JwtToken := token.JwtToken{
		Token: tokenString,
	}

	claims, err := JwtToken.Verify()
	if err != nil {
		response.ShowErr(c, http.StatusUnauthorized, "Invalid token", nil)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"user":   claims,
	})
}
