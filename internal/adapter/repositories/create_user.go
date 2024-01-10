package repositories

import (
	"context"
	"fmt"

	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/repositories/model"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"gorm.io/gorm"
)

type userCreator repo

var _ interfaces.RCreateUser = (*userCreator)(nil)

func NewUserCreator(
	db *gorm.DB,
) *userCreator {
	return &userCreator{
		DB: db,
	}
}

func (r *userCreator) Do(ctx context.Context, in entity.User) error {
	var (
		userModel model.User
	)

	if err := userModel.FromEntity(in); err != nil {
		return err
	}

	fmt.Println(userModel)

	return r.DB.
		WithContext(ctx).
		Omit("Swipes").
		Debug().
		Create(&userModel).Error
}
