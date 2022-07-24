package model

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
)

type user_role string

const (
	admin  user_role = "admin"
	waiter           = "waiter"
)

type User struct {
	Username          string    `json:"username"`
	Password          string    `json:"password,omitempty"`
	EncryptedPassword string    `json:"-"`
	Role              user_role `json:"role"`
}

func (u *User) EncryptPassword() error {
	sha := sha256.New()
	if _, err := sha.Write([]byte(u.Password)); err != nil {
		return err
	}
	fmt.Println(sha.Sum(nil))
	u.EncryptedPassword = base64.StdEncoding.EncodeToString(sha.Sum(nil))
	return nil
}

func (u *User) CompareWithHash(password string) (error, bool) {
	sha := sha256.New()
	if _, err := sha.Write([]byte(password)); err != nil {
		return err, false
	}
	a := base64.StdEncoding.EncodeToString(sha.Sum(nil))
	fmt.Println(a, u)
	if u.EncryptedPassword == a {
		return nil, true
	}
	return errors.New("password incorrect"), false
}

func (u *User) Sanitize() {
	u.Password = ""
}
