package db

import (
	"fmt"
	"log"

	"github.com/yafiesetyo/dating-apps-srv/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(cfg config.Config) *gorm.DB {
	dsn := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(fmt.Sprintf(
		dsn, cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.DbName, cfg.Db.Port)), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init DB, err: %v", err)
	}

	return db
}
