package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPasswword(password string) (string,error){
	hashePassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		return "", fmt.Errorf("failed to hash password: %w",err)
	}
	return string(hashePassword),nil
}

func CheckPassword(password string, hashedPassword string) error{
	// fmt.Printf("%s",hashedPassword)\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
