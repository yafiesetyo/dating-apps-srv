package repositories

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/repositories/model"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"gorm.io/gorm"
)

type userSwipeCreator repo

var _ interfaces.RCreateUserSwipe = (*userSwipeCreator)(nil)

func NewUserSwipeCreator(
	db *gorm.DB,
) *userSwipeCreator {
	return &userSwipeCreator{
		DB: db,
	}
}

func (r *userSwipeCreator) Do(ctx context.Context, in entity.UserSwipe) error {
	var (
		userSwipeModel model.UserSwipe
	)
	userSwipeModel.FromEntity(in)

	return r.DB.WithContext(ctx).
		Omit("LikedUser").
		Create(&userSwipeModel).
		Error
}
