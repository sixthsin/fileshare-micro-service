package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Username string `gorm:"type:varchar(50);uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
