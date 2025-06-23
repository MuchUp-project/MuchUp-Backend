package database 

import (
	"fmt"
	"gorm.io/gorm"
	"MuchUp/backend/config"
	"gorm.io/driver/postgres"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
        cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
    return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}