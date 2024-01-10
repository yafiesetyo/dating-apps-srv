package entity

import (
	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
)

type User struct {
	ID             uint
	Name           string
	Gender         constants.Gender
	Username       string
	Password       string
	ImageUrl       []string
	DOB            string
	POB            string
	Religion       constants.Religion
	Description    string
	Hobby          string
	VerifiedUser   bool
	UnlimitedSwipe bool
}
