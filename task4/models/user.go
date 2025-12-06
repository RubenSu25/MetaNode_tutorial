package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"uniqueIndex;size:50;not null" json:"username"`
	Password string `gorm:"size:255;not null" json:"password"`
	Email    string `gorm:"size:100;not null" json:"email"`

	// 关系定义
	Posts []Post `gorm:"foreignKey:UserID" json:"posts"`
}
