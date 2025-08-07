package database

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	Title   string `json:"title"`
	Extract string `json:"extract"`
	Query   string `json:"query" gorm:"index"`
}
