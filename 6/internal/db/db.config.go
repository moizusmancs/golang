package internal

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgre() *gorm.DB{

	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5440 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("❌database initialization failed: %v", err)
	}

	log.Println("✅ DB connected with server")
	return db;

}