package middleware

import (
	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("c:wOAy{x%V^]I.`u=BL")

type UserClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func JwtCreate(userID int, expiredAt int64) string {
	claims := UserClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expiredAt,
			Issuer:    "omg",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(mySigningKey)

	return ss
}
