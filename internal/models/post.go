package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ProfileID uint
	Creator   string
	Text      string
	Likes     int `gorm:"default:0" json:"-"`
	Dislikes  int `gorm:"default:0" json:"-"`

	Comments []Comment `gorm:"default:[], foreignKey:PostID" json:"-"`
}
