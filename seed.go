package main

import (
	"fmt"

	"github.com/Mohammad-Alipour/menu-api/config"
	"github.com/Mohammad-Alipour/menu-api/models"
)

func Seed() {
	// Categories
	categories := []models.Category{
		{Name: "نوشیدنی‌ها", Description: "انواع نوشیدنی سرد و گرم", Active: true},
		{Name: "غذاها", Description: "غذاهای ایرانی و فرنگی", Active: true},
	}

	for _, cat := range categories {
		config.DB.FirstOrCreate(&cat, models.Category{Name: cat.Name})
	}

	// Items
	items := []models.Item{
		{CategoryID: 1, Name: "قهوه اسپرسو", Description: "قهوه تلخ و غلیظ", Price: 50000, Active: true},
		{CategoryID: 1, Name: "چای ایرانی", Description: "چای سیاه خوش‌طعم", Price: 20000, Active: true},
		{CategoryID: 2, Name: "پیتزا مخلوط", Description: "پیتزا با گوشت و سبزیجات", Price: 150000, Active: true},
		{CategoryID: 2, Name: "قرمه سبزی", Description: "غذای سنتی ایرانی", Price: 120000, Active: true},
	}

	for _, item := range items {
		config.DB.FirstOrCreate(&item, models.Item{Name: item.Name})
	}

	fmt.Println("✅ Initial data inserted successfully")
}
