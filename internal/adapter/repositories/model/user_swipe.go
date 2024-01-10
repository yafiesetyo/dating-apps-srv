package model

import (
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"gorm.io/gorm"
)

type (
	UserSwipe struct {
		gorm.Model

		UserID      uint  `gorm:"column:user_id"`
		LikedUserID *uint `gorm:"column:liked_user_id"`

		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

		LikedUser User `gorm:"foreignKey:LikedUserID;references:ID"`
	}
)

func (UserSwipe) TableName() string {
	return "user_swipes"
}

func (u *UserSwipe) FromEntity(in entity.UserSwipe) {
	*u = UserSwipe{
		UserID:      in.UserID,
		LikedUserID: in.LikedUserID,
	}
}
