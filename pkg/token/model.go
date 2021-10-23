package token

type JwtToken struct {
	Token string `json:"token"`
}

const jwtSecretKey = "563343c3735166379e23369a1e4a2562"
const signingMethod = "HS256"
