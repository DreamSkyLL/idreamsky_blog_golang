package users

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password string
	Name     string
	Email    *string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hash)
	return
}

func (u *User) CheckPassword(pwd string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		fmt.Println("Wrong PWD")
	} else {
		fmt.Println("Right PWD")
	}
	return
}
