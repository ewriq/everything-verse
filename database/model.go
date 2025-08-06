package database

import "gorm.io/gorm"

type Data struct {
	gorm.Model
	Title   string 
	Extract string 
	Query   string 
}

