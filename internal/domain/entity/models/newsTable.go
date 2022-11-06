package models

import "github.com/jinzhu/gorm"

type NewsTable struct {
	gorm.Model
	UserId uint   `json:"user_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Active bool   `json:"active"`
	Text   string `json:"text"`
}
