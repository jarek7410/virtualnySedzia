package model

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title   string
	Content string
}
