package db

import (
	"fmt"
	"log"
	"os"

	"backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	user := getEnv("DB_USER", "root")
	pass := getEnv("DB_PASS", "")
	host := getEnv("DB_HOST", "127.0.0.1")
	port := getEnv("DB_PORT", "3306")
	name := getEnv("DB_NAME", "elearn")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Gagal koneksi DB:", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Student{}, &models.Course{}, &models.StudentCourse{})
	log.Println("✅ DB connected & migrated")
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
