package database

import (
	"log"

	"moneyManagerAPI/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/money_manager?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("failed to connect database ", err.Error())
	}

	db.AutoMigrate(models.Category{})
	db.AutoMigrate(models.Transaction{})
	db.AutoMigrate(models.Home{})
	db.AutoMigrate(models.User{})
	return db
}
