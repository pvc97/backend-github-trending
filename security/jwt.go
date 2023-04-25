package security

import (
	"backend-github-trending/model"
	"time"

	"github.com/golang-jwt/jwt"
)

const JWT_KEY = "d8a5a38349f92995c1d5f49f36cd987efe6d1be9067dcef1bdeca4411d433d38e98d275d3fbcc11af7d5d2d40d0e19afef24a9bd3e3ba174aaea3300b772b382"

func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(JWT_KEY))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
