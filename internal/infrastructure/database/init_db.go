package database 

import (
	"gorm.io/gorm"
	"github.com/MuchUp/backend/internal/domain/entity"
)

func InitDB(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&entity.User{})
	db.AutoMigrate(&entity.ChatGroup{})
	db.AutoMigrate(&entity.Message{})
	return db
}

