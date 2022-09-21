package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username,omitempty" validate:"required"`
	Password string `json:"password,omitempty" validate:"required"`
}
