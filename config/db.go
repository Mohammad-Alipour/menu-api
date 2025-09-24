package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	//DatabaseConnectionInfo
	dsn := "host=localhost user=postgres password=yourpassword dbname=menu_api port=5432 sslmode=disable TimeZone=Asia/Tehran"

	//ConnectDatabase
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ اتصال به دیتابیس ناموفق بود: ", err)
	}

	DB = db
	fmt.Println("✅ اتصال به دیتابیس برقرار شد")
}
