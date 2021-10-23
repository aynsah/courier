package middleware

import (
	"courier/pkg/response"
	"courier/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")

	JwtToken := token.JwtToken{
		Token: tokenString,
	}

	claims, err := JwtToken.Verify()
	if err != nil {
		response.ShowErr(c, http.StatusUnauthorized, "Invalid token", nil)
		c.Abort()
		return
	}

	c.Set("claims", claims)
}
