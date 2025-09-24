package config

import (
	"fmt"
	"log"

	"github.com/Mohammad-Alipour/menu-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=1234 dbname=menu_api port=5432 sslmode=disable TimeZone=Asia/Tehran"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to the database: ", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Item{},
		&models.ItemImage{},
	)
	if err != nil {
		log.Fatal("❌ Migration failed: ", err)
	}

	_ = db
	fmt.Println("✅ Database connection established and migration completed")
}
