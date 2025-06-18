package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ProfileID uint
	PostID  uint
	Creator string
	Text    string
}
