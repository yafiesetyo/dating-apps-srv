package request

import (
	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
)

type Register struct {
	Name        string   `json:"name" validate:"required,alpha"`
	Gender      string   `json:"gender" validate:"required,oneof=male female"`
	Username    string   `json:"username" validate:"required,alphanum"`
	Password    string   `json:"password" validate:"required,gte=8"`
	ImageUrl    []string `json:"image_url" validate:"required,dive,required,http_url"`
	DOB         string   `json:"dob" validate:"required"`
	POB         string   `json:"pob" validate:"required"`
	Religion    string   `json:"religion" validate:"required,oneof=islam kristen katolik hindu buddha konghuchu lainnya"`
	Description string   `json:"description" validate:"required"`
	Hobby       string   `json:"hobby" validate:"required"`
}

func (r Register) ToEntity() entity.User {
	return entity.User{
		Name:        r.Name,
		Gender:      constants.Gender(r.Gender),
		Username:    r.Username,
		Password:    r.Password,
		ImageUrl:    r.ImageUrl,
		DOB:         r.DOB,
		POB:         r.POB,
		Religion:    constants.Religion(r.Religion),
		Description: r.Description,
		Hobby:       r.Description,
	}
}
