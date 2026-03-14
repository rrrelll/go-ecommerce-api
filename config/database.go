package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {

	dsn := "root:akbarfarel802@tcp(127.0.0.1:3306)/go_shop?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database Connected")

	return db
}
