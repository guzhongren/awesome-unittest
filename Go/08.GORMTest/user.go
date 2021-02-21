package main

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

// CreateUser 创建一个用户
func (u User) CreateUser() {
	db.Create(&u)
}
