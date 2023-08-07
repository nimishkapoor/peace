package polling

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("akakakakakakakak")

func GenerateJWT(username string, tenant string, role string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["tenant"] = tenant
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 300).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(token string) (jwt.MapClaims, error) {
	token_c, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, _ := token_c.Claims.(jwt.MapClaims)
	return claims, nil
}
