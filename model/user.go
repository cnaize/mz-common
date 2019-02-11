package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Base
	Username  string     `json:"username" gorm:"unique_index" form:"username" binding:"required"`
	Password1 string     `json:"-" sql:"-" form:"password1" binding:"required"`
	Password2 string     `json:"-" sql:"-" form:"password2" binding:"required"`
	PassHash  string     `json:"-" form:"-"`
	DeletedAt *time.Time `json:"-"`
	Error     *Error     `json:"error,omitempty" sql:"-" form:"error"`
}

type Token struct {
	jwt.StandardClaims
	Username string
}
