package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Base
	Username  string     `json:"username" gorm:"unique_index" form:"username" binding:"required"`
	Token     string     `json:"token" sql:"-"`
	Password  *string    `json:"password,omitempty" sql:"-" form:"password" binding:"required"`
	Password1 *string    `json:"password1,omitempty" sql:"-" form:"password1"`
	PassHash  string     `json:"-" form:"-"`
	DeletedAt *time.Time `json:"-"`
	Error     *Error     `json:"error,omitempty" sql:"-" form:"error"`
}

type Token struct {
	jwt.StandardClaims
	Username string
}
