package database

type Data struct {
	ID      uint   `gorm:"not null;primaryKey" json:"id"`
	Title   string `gorm:"not null" json:"title"`
	Extract string `gorm:"not null" json:"extract"`
	Source  string `gorm:"not null;uniqueIndex" json:"source"`
	Url     string `gorm:"not null" json:"url"`
}

type Extract struct {
	ID      uint   `gorm:"not null;primaryKey"`
	Extract string `gorm:"not null" json:"data"`
}
