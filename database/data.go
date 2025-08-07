package database

import "fmt"

func Insert(query, extract, title string) error {
	data := Data{
		Query:   query,
		Extract: extract,
		Title:   title,
	}

	result := db.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAll() ([]Data, error) {
	var data []Data
	result := db.Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func Get(query string) ([]Data, error) {
	var data []Data
	result := db.Where("query = ?", query).Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func SearchFTS(keyword string) ([]Data, error) {
	var results []Data

	if keyword == "" {
		return results, nil
	}

	likePattern := "%" + keyword + "%"
	result := db.Where("title LIKE ? OR extract LIKE ? OR query LIKE ?",
		likePattern, likePattern, likePattern).Find(&results)

	return results, result.Error
}

func Exists(query string) bool {
	var count int64
	err := db.Model(&Data{}).Where("query = ?", query).Count(&count).Error
	if err != nil {
		fmt.Println("DB error:", err)
		return false
	}
	return count > 0
}
