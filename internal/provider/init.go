package provider

import (
	"TWclone/internal/config"
	"TWclone/internal/database"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitGlobal(cfg *config.Config) {
	var err error
	db, err = database.InitGorm(cfg)
	if err != nil {
		panic(err)
	}
}
