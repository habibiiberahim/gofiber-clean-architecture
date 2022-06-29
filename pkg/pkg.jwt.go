package pkg

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

type MetaToken struct {
	ID        string
	Email     string
	ExpiredAt time.Time
	Authorization bool
}

type AccessToken struct{
	Claims MetaToken
}

func Sign(Data map[string]interface{}, SecretPublicKeyEnvName string, ExpiredAt time.Duration) (string , error){
	expiredAt := time.Now().Add(time.Duration(time.Minute)*ExpiredAt).Unix()

	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	claims := jwt.MapClaims{}
	claims["exp"] = expiredAt
	claims["authorization"] = true

	for i, v := range Data {
		claims[i] = v
	}

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err:= to.SignedString([]byte(jwtSecretKey))
	
	if err != nil {
		logrus.Error(err.Error())
		return accessToken, err
	}

	return accessToken, nil
}

func VerifyTokenHeader(ctx *fiber.Ctx, SecretPublicKeyEnvName string) (*jwt.Token, error){
	tokenHeader := ctx.Get("Authorization")
	accessToken := strings.Split(tokenHeader, "Bearer")[1] 
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(strings.Trim(accessToken," "),func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey),nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
	
}

func VerifyToken(accessToken, SecretPublicKeyEnvName string) (*jwt.Token, error){
	jwtSecretKey := GodotEnv(SecretPublicKeyEnvName)

	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey),nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}

func DecodeToken(accessToken *jwt.Token) AccessToken{
	var token AccessToken
	stringify, _ := json.Marshal(&accessToken)
	json.Unmarshal([]byte(stringify),&token)

	return token
}