package config

import (
	"fmt"
	"github.com/Mohammad-Alipour/menu-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var _ *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=menu_api port=5432 sslmode=disable TimeZone=Asia/Tehran"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ اتصال به دیتابیس ناموفق بود: ", err)
	}

	// Migration
	db.AutoMigrate(&models.User{})

	_ = db
	fmt.Println("✅ اتصال به دیتابیس برقرار شد و Migration انجام شد")
}
