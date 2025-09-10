package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type MyCustomClaims struct {
	ID   int
	Name string
	Role string
	jwt.RegisteredClaims
}

var my_Key = "my_secret_key"

func GenerateToken(u_Id int, u_Name, u_Role string) (string, error) {
	claims := MyCustomClaims{
		ID:   u_Id,
		Name: u_Name,
		Role: u_Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "MyTaskManager",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(my_Key))
}

func VerifyJWT(tokenString string)(string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(my_Key), nil
	})

	if err != nil || !token.Valid {
		return "", err
	}

     claims, ok := token.Claims.(*MyCustomClaims)

	if !ok {
		return "",fmt.Errorf("could not parse claims")
	}
	return claims.Role, nil
}
