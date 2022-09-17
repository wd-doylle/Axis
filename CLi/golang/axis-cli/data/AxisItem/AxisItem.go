package AxisItem

import (
	"axis-cli/data"
	"axis-cli/data/AxisUniverse"
	"time"

	"gorm.io/gorm"
)

type AxisItem struct {
	gorm.Model
	ID          uint64 `gorm:"primaryKey"`
	Title       string
	Description string
	TimeDue     time.Time
	Priority    uint8
	UniverseID  uint64
	Universe    AxisUniverse.AxisUniverse `gorm:"foreignKey:UniverseID"`
	Status      string
}

func Create(title string, description string, priority uint8, timeDue time.Time, universeName string) error {
	// Create
	var universe AxisUniverse.AxisUniverse
	result := data.Db.First(&universe, "name = ?", universeName)
	if result.Error != nil {
		return result.Error
	}
	result = data.Db.Create(&AxisItem{Title: title, Description: description, Priority: priority, TimeDue: timeDue, Universe: universe})
	return result.Error
}

func List() ([]AxisItem, error) {
	// Read
	var items []AxisItem
	result := data.Db.Model(&AxisItem{}).Limit(10).Find(&items)
	return items, result.Error
}

func ListToMap() ([]map[string]interface{}, error) {
	// Read
	var itemMaps []map[string]interface{}
	result := data.Db.Model(&AxisItem{}).Limit(10).Find(&itemMaps)
	return itemMaps, result.Error
}

// func Find(name string) (AxisUniverse, error) {
// 	// Read
// 	var universe AxisUniverse
// 	result := data.Db.First(&universe, "name = ?", name)
// 	return universe, result.Error
// }

// func Update(name string, newName string, newDescription string) error {
// 	var universe AxisUniverse
// 	result := data.Db.First(&universe, "name = ?", name)
// 	if newName != "" {
// 		universe.Name = newName
// 	}
// 	if newDescription != "" {
// 		universe.Description = newDescription
// 	}
// 	// data.Db.Model(&universe).Updates(universe)
// 	data.Db.Save(&universe)
// 	return result.Error
// }

// func Delete(name string) error {
// 	var universe AxisUniverse
// 	result := data.Db.First(&universe, "name = ?", name)
// 	data.Db.Delete(&universe)
// 	return result.Error
// }

// func (u *AxisUniverse) BeforeDelete(tx *gorm.DB) (err error) {
// 	if u.Role == "admin" {
// 		return errors.New("admin user not allowed to delete")
// 	}
// 	return
// }

func init() {
	// Migrate the schema
	data.Db.AutoMigrate(&AxisItem{})
}
