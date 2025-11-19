package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"index"`
	Password string `json:"password"`
	NickName string `json:"nickname"`
}

func NewUser(email string, password string, nickName string) *User {
	user := &User{
		Email:    email,
		Password: password,
		NickName: nickName,
	}
	return user
}
