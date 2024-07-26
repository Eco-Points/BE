package utils

import (
	"eco_points/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtUtilityInterface interface {
	GenereteJwt(id uint) (string, error)
	DecodToken(token *jwt.Token) uint
}

type JwtUtility struct{}

func NewJwtUtility() JwtUtilityInterface {
	return &JwtUtility{}
}

func (ju *JwtUtility) GenereteJwt(id uint) (string, error) {

	jwtKey := config.ImportSetting().JWTSecret
	data := jwt.MapClaims{}

	data["id"] = id
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Minute * 45).Unix()

	processToken := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	result, err := processToken.SignedString([]byte(jwtKey))

	if err != nil {
		return "", err
	}

	return result, nil
}

func (ju *JwtUtility) DecodToken(token *jwt.Token) uint {
	var result uint

	claim := token.Claims.(jwt.MapClaims)

	for _, val := range claim {
		fmt.Println(val)
	}

	if value, found := claim["id"]; found {
		result = value.(uint)
	}

	return result
}
