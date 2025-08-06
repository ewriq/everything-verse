package database

import (
	"log"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error



func init() {
	db, err = gorm.Open(sqlite.Open("database/data.db?_journal_mode=WAL&_busy_timeout=5000"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(1)
	}

	log.Println("✅ GORM ile Sqlite bağlantısı kuruldu.")
	db.AutoMigrate(&Data{})
}
