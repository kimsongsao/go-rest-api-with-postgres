package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string
	Body   string
	Likes  int
	Draft  bool
	Author string
	// UserID uint `gorm:"user_id"`
}
