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

	var db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(`❌ Failed to connect to the database: `, err)
	}

	// Migration
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	_ = db
	fmt.Println("✅ Database connection established and migration completed")
}
