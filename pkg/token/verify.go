package token

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func (j JwtToken) Verify() (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(j.Token, claims, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod(signingMethod) != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecretKey), nil
	})

	if token == nil || err != nil {
		return nil, err
	}

	return claims, nil
}
