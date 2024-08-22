package model

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDb() {
	dir := "./db/"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	db, _ = gorm.Open(sqlite.Open(dir+"image.db"), &gorm.Config{})
	db.AutoMigrate(&ImageInfo{})
}
