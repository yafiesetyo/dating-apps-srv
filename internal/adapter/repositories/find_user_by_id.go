package repositories

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/repositories/model"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"gorm.io/gorm"
)

type userFinderById repo

var _ interfaces.RFindUserByID = (*userFinderById)(nil)

func NewUserFinderById(
	db *gorm.DB,
) *userFinderById {
	return &userFinderById{
		DB: db,
	}
}

func (r *userFinderById) Do(ctx context.Context, id uint) (entity.User, error) {
	var (
		userModel model.User
	)

	err := r.DB.
		WithContext(ctx).
		Where("id=?", id).
		First(&userModel).
		Error
	if err != nil {
		return entity.User{}, err
	}

	return userModel.ToEntity(), nil
}
