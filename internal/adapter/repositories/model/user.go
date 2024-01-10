package model

import (
	"time"

	"github.com/lib/pq"
	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model

		Name           string         `gorm:"column:name"`
		Gender         string         `gorm:"column:gender"`
		Username       string         `gorm:"column:username"`
		Password       string         `gorm:"column:password"`
		ImageUrl       pq.StringArray `gorm:"type:text[],column:image_url"`
		DOB            time.Time      `gorm:"column:dob"`
		POB            string         `gorm:"column:pob"`
		Religion       string         `gorm:"column:religion"`
		Description    string         `gorm:"column:description"`
		Hobby          string         `gorm:"column:hobby"`
		VerifiedUser   bool           `gorm:"column:verified_user"`
		UnlimitedSwipe bool           `gorm:"column:unlimited_swipe"`

		DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

		Swipes []UserSwipe `gorm:"foreignKey:UserID;references:ID"`
	}
)

func (User) TableName() string {
	return "users"
}

func (u *User) FromEntity(in entity.User) error {
	dob, err := time.Parse(constants.DoBFormat, in.DOB)
	if err != nil {
		return err
	}

	*u = User{
		Model: gorm.Model{
			ID: in.ID,
		},
		Name:           in.Name,
		Gender:         string(in.Gender),
		Username:       in.Username,
		Password:       in.Password,
		ImageUrl:       in.ImageUrl,
		VerifiedUser:   in.VerifiedUser,
		UnlimitedSwipe: in.UnlimitedSwipe,
		DOB:            dob,
		POB:            in.POB,
		Religion:       string(in.Religion),
		Description:    in.Description,
		Hobby:          in.Hobby,
	}

	return nil
}

func (u *User) ToEntity() entity.User {
	return entity.User{
		ID:             u.ID,
		Name:           u.Name,
		Gender:         constants.Gender(u.Gender),
		Username:       u.Username,
		Password:       u.Password,
		ImageUrl:       u.ImageUrl,
		VerifiedUser:   u.VerifiedUser,
		UnlimitedSwipe: u.UnlimitedSwipe,
		DOB:            u.DOB.Format(constants.DoBFormat),
		POB:            u.POB,
		Religion:       constants.Religion(u.Religion),
		Description:    u.Description,
		Hobby:          u.Hobby,
	}
}
