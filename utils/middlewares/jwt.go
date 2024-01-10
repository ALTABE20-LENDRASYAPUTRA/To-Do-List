package middlewares

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(idUser uint) (string, error) {
	var claim = jwt.MapClaims{}
	claim["id"] = idUser
	claim["iat"] = time.Now().UnixMilli()
	claim["exp"] = time.Now().Add(time.Hour * 1).UnixMilli()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	// token := jwt.New(jwt.SigningMethodHS256)
	strToken, err := token.SignedString([]byte("K3YYKijS879!!"))
	if err != nil {
		return "", err
	}

	return strToken, nil
}

func ExtractTokenUserId(t *jwt.Token) (uint, error) {
	var userID uint

	if t.Valid {
		var tokenClaims = t.Claims.(jwt.MapClaims)
		userID = uint(tokenClaims["id"].(float64))

		return userID, nil
	}

	return 0, errors.New("token tidak valid")
}