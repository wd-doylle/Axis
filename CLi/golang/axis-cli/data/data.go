package data

import (
	"log"
	"os"
	"path/filepath"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	axisDir := filepath.Join(homeDir, ".axis")
	os.MkdirAll(axisDir, os.ModePerm)
	Db, err = gorm.Open(sqlite.Open(filepath.Join(axisDir, "axis.db")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}
