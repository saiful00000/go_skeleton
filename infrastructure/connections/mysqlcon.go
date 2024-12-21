package connections

import (
	"fmt"
	"go_clean/core/models"
	"go_clean/infrastructure/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func connect() {
	config := configs.LocalConfig

	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUser, config.DbPass, config.DbIp, config.DbName)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Error while connecting to DB")
		panic(err)
	}

	fmt.Println("DB connected.")

	db = d

}

func migrate() {
	db.Migrator().AutoMigrate(&models.Useer{})
}

func GetDB() *gorm.DB{
	if db == nil {
		connect()
	}

	migrate()

	return db
}
