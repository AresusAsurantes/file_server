package data

import (
	"crypto/rand"
	"fmt"
	"log"
)

type User struct {
	Id       string
	UserName string
	PassWord string
	Role     string
}

var UuidLength = 10

func NewUser(username string, password string) *User {
	return &User{
		Id:       uuid(),
		UserName: username,
		PassWord: password,
		Role:     "USER",
	}
}

func uuid() string {
	b := make([]byte, UuidLength)
	_, err := rand.Read(b)
	if err != nil {
		log.Panic(err)
	}

	return fmt.Sprintf("%x", b)
}
