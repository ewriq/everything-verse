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
	db, err = gorm.Open(
		sqlite.Open("database/data.db?_journal_mode=WAL&_busy_timeout=5000"),
		&gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		},
	)

	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	sqlDB, err := db.DB()
	if err == nil {
		sqlDB.SetMaxOpenConns(1)
	}

	log.Println("Connection established.")

	err = db.AutoMigrate(&Data{}, &Extract{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	err = db.Exec(`
		CREATE VIRTUAL TABLE IF NOT EXISTS data_fts USING fts5(
			title,
			extract
		)
	`).Error

	if err != nil {
		log.Fatalf("FTS5 creation failed: %v", err)
	}

	log.Println("FTS5 ready.")
}