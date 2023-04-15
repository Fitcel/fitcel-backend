package configuration

import (
	"fitcel-backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect() *gorm.DB {
	dsn := "root@tcp(127.0.0.1:3306)/diet_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil
	}
	db.AutoMigrate(&models.Celebrities{})
	db.AutoMigrate(new(models.Diet), new(models.Food), new(models.Meal))
	db.AutoMigrate(new(models.User))
	return db
}
