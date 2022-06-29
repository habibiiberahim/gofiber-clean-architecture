package pkg

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
			logrus.Fatal(err.Error())
	}
	return string(result)
}

func ComparePassword(plainPassword string, hashPassword string) error  {
	plainPw := []byte(plainPassword) 
	hashPw := []byte(hashPassword) 
	err := bcrypt.CompareHashAndPassword(hashPw, plainPw)
	return err
}