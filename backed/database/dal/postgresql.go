package dal

import (
	"fmt"

	"Backed/config"
	"Backed/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgreSQL *gorm.DB

func InitMySQL(cfg config.DBConfig) error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	PostgreSQL = db
	return PostgreSQL.AutoMigrate(&model.User{}, &model.Article{})
}
