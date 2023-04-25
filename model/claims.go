package model

import "github.com/golang-jwt/jwt"

// Claims means the data in the token
type JwtCustomClaims struct {
	UserId string `json:"userId"`
	Role   string `json:"role"`
	jwt.StandardClaims
}
