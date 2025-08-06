package database

import (
	"log"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	db, err = gorm.Open(sqlite.Open("database/data.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	log.Println("✅ GORM ile Sqlite bağlantısı kuruldu.")
	db.AutoMigrate(&Data{})
	db.Exec("PRAGMA journal_mode = WAL;")

}

