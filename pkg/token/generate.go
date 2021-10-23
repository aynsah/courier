package token

import "github.com/dgrijalva/jwt-go"

func Generate(UUID string, MSISDN string, username string) (JwtToken, error) {
	JwtToken := JwtToken{}
	sign := jwt.New(jwt.GetSigningMethod(signingMethod))
	claims := sign.Claims.(jwt.MapClaims)

	claims["uuid"] = UUID
	claims["msisdn"] = MSISDN
	claims["username"] = username

	var err error

	JwtToken.Token, err = sign.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return JwtToken, err
	}

	return JwtToken, nil

}
