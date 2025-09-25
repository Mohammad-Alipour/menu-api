package config

import (
	"fmt"
	"log"

	"github.com/Mohammad-Alipour/menu-api/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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

	DB = db
	fmt.Println("✅ Database connection established and migration completed")

	var count int64
	configUser := models.User{
		Name:  "Admin",
		Email: "admin@cafe.com",
		Role:  "admin",
	}
	DB.Model(&models.User{}).Where("email = ?", configUser.Email).Count(&count)
	if count == 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("1234"), bcrypt.DefaultCost)
		configUser.PasswordHash = string(hashedPassword)
		DB.Create(&configUser)
		fmt.Println("✅ Default admin user created (email: admin@cafe.com, password: 1234)")
	}
}
