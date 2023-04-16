package configuration

import (
	"fitcel-backend/models"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func dbConnect(runmode string) *gorm.DB {
	mysqlConf := viper.GetStringMapString(runmode + ".mysql")
	dbParams := "?charset=utf8mb4&parseTime=True&loc=Local"
	if runmode == "prod" {
		dbParams += "&tls=true"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s%s", mysqlConf["user"], mysqlConf["password"], mysqlConf["host"], mysqlConf["database"], dbParams)
	fmt.Println(dsn)
	// dsn := "root@tcp(127.0.0.1:3306)/diet_service?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return nil
	}
	db.AutoMigrate(&models.Celebrity{})
	db.AutoMigrate(new(models.Diet), new(models.Food), new(models.Meal))
	db.AutoMigrate(new(models.User))
	return db
}
