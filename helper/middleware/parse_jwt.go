package middleware

import (
	"BE/model"
	"encoding/json"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

func ClaimJwt(tokenString string) (model.User, error) {
	user := model.User{}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return user, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userClaimMarshaled, _ := json.Marshal(claims)
		if err := json.Unmarshal(userClaimMarshaled, &user); err != nil {
			return user, err
		}
	}
	return user, nil
}
