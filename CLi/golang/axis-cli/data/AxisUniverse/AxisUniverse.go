package AxisUniverse

import (
	"axis-cli/data"

	"gorm.io/gorm"
)

type AxisUniverse struct {
	gorm.Model
	Name        string
	Description string
}

func Create(name string, description string) {
	// Create
	data.Db.Create(&AxisUniverse{Name: name, Description: description})
}

func List() []AxisUniverse {
	// Read
	var universes []AxisUniverse
	data.Db.Limit(10).Find(&universes)
	return universes
	// db.First(&universe, 1) // find product with integer primary key
	// db.First(&universe, "code = ?", "D42") // find product with code D42
}

func Update() {
	// var universe AxisUniverse
	// db.First(&universe, 1) // find product with integer primary key
	// // Update - update product's price to 200
	// db.Model(&universe).Update("Price", 200)
	// // Update - update multiple fields
	// db.Model(&universe).Updates(AxisUniverse{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&universe).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}

func Delte() {
	// var universe AxisUniverse
	// db.First(&universe, 1) // find product with integer primary key
	// // Delete - delete product
	// db.Delete(&universe, 1)
}

func init() {
	// Migrate the schema
	data.Db.AutoMigrate(&AxisUniverse{})
}
