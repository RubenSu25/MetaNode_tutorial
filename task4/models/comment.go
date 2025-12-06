package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`

	// 关系定义
	UserID uint `gorm:"index;not null" json:"user_id"`
	User   User `gorm:"foreignKey:UserID" json:"author"`

	PostID uint `gorm:"index;not null" json:"post_id"`
	Post   Post `gorm:"foreignKey:PostID" json:"post"`
}
