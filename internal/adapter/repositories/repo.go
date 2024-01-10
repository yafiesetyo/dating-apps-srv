package repositories

import "gorm.io/gorm"

type repo struct {
	DB *gorm.DB
}
