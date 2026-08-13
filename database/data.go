package database

import (
	"fmt"

	"gorm.io/gorm"
)

func Insert(source, extract, title, url string) error {
	return db.Transaction(func(tx *gorm.DB) error {

		insertExtract := Extract{
			Extract: extract,
		}

		if err := tx.Create(&insertExtract).Error; err != nil {
			return err
		}

		data := Data{
			Extract: fmt.Sprintf("%d", insertExtract.ID),
			Title:   title,
			Source:  source,
			Url:     url,
		}

		if err := tx.Create(&data).Error; err != nil {
			return err
		}

		if err := tx.Exec(`
            INSERT INTO data_fts(rowid, title, extract)
            VALUES (?, ?, ?)
        `, data.ID, data.Title, extract).Error; err != nil {
			return err
		}

		return nil
	})
}

func GetAll() (int64, error) {
	var count int64

	if err := db.Model(&Data{}).Count(&count).Error; err != nil {
		return  0, err
	}

	return  count, nil
}

func SearchFTS(keyword string) ([]Data, error) {
	var results []Data

	if keyword == "" {
		return results, nil
	}

	err := db.Raw(`
		SELECT
			data.id,
			data.title,
			extracts.extract AS extract,
			data.source,
			data.url
		FROM data
		JOIN data_fts
			ON data.id = data_fts.rowid
		JOIN extracts
			ON extracts.id = CAST(data.extract AS INTEGER)
		WHERE data_fts MATCH ?
		ORDER BY bm25(data_fts)
	`, keyword).Scan(&results).Error

	return results, err
}

func Exists(source string) bool {
	var count int64
	err := db.Model(&Data{}).Where("source = ?", source).Count(&count).Error
	if err != nil {
		fmt.Println("DB error:", err)
		return false
	}
	return count > 0
}

