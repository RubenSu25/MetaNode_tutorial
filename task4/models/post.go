package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title   string `gorm:"size:200;not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`

	// 关系定义
	UserID   uint      `gorm:"index;not null" json:"user_id"`
	User     User      `gorm:"foreignKey:UserID" json:"author,omitempty"`
	Comments []Comment `gorm:"foreignKey:PostID" json:"comments,omitempty"`
}
