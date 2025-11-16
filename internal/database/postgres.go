package database

import (
	"fmt"
	"time"

	"TWclone/internal/config"
	"TWclone/internal/entity"
	"TWclone/internal/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGorm(cfg *config.Config) (*gorm.DB, error) {
	dbCfg := cfg.Database

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Jakarta",
		dbCfg.Host,
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.DbName,
		dbCfg.Port,
		dbCfg.Sslmode,
	)

	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log.Fatalf("failed to connect to gorm database: %v", err)
		return nil, err
	}

	// Set package-level DB so repositories using database.DB won't be nil
	DB = gdb

	sqlDB, err := gdb.DB()
	if err != nil {
		logger.Log.Fatalf("failed to get sql DB from gorm: %v", err)
		return nil, err
	}
	sqlDB.SetMaxIdleConns(dbCfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbCfg.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(dbCfg.MaxConnLifetime) * time.Minute)

	if err := gdb.AutoMigrate(&entity.User{}); err != nil {
		logger.Log.Fatalf("failed to run automigrate: %v", err)
		return nil, err
	}

	return gdb, nil
}
