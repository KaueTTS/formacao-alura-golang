package postgresdb

import (
	"fmt"
	"log"
	"time"

	"go_gin_criando_api_rest_com_simplicidade/src/config/env"
	"go_gin_criando_api_rest_com_simplicidade/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg env.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Pass, cfg.DB.Name, cfg.DB.SSLMode, cfg.DB.Timezone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("sql DB error: %v", err)
	}
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	if err := db.AutoMigrate(&models.Student{}); err != nil {
		log.Fatalf("auto-migrate error: %v", err)
	}

	DB = db
	return db
}

func Close() {
	if DB == nil {
		return
	}
	if sqlDB, err := DB.DB(); err == nil {
		_ = sqlDB.Close()
	}
}
