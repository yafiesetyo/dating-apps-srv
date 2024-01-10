package repositories

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/repositories/model"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"gorm.io/gorm"
)

type userFinderByUsername repo

var _ interfaces.RFindUserByUsername = (*userFinderByUsername)(nil)

func NewUserFinderByUsername(
	db *gorm.DB,
) *userFinderByUsername {
	return &userFinderByUsername{
		DB: db,
	}
}

func (r *userFinderByUsername) Do(ctx context.Context, username string) (entity.User, error) {
	var (
		userModel model.User
	)

	err := r.DB.
		WithContext(ctx).
		Where("username=?", username).
		First(&userModel).
		Error
	if err != nil {
		return entity.User{}, err
	}

	return userModel.ToEntity(), nil
}
