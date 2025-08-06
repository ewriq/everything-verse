package database

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
