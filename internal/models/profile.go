package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=8"`
	Username       string `json:"username" validate:"required,min=4"`
	Bio            string `json:"bio" validate:"required"`
	Github         string `json:"github" validate:"required"`
	FollowersCount int    `gorm:"default:0" json:"-"`

	Posts     []Post         `gorm:"foreignKey:ProfileID" json:"-"`
	Followers pq.StringArray `gorm:"type:text[]" json:"-"`
	Following pq.StringArray `gorm:"type:text[]" json:"-"`
}
