package repositories

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/repositories/model"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"gorm.io/gorm"
)

type userProfileFinder repo

var _ interfaces.RFindUserProfile = (*userProfileFinder)(nil)

func NewUserProfileFinder(
	db *gorm.DB,
) *userProfileFinder {
	return &userProfileFinder{
		DB: db,
	}
}

func (r *userProfileFinder) Do(ctx context.Context, id []uint) ([]entity.User, error) {
	var (
		userModels []model.User
		userSwipes []model.UserSwipe
	)

	err := r.DB.WithContext(ctx).
		Raw(`select us.* from user_swipes us where user_id=?`, id[0]).
		Scan(&userSwipes).
		Error
	if err != nil {
		return []entity.User{}, err
	}

	skippedIds := []uint{}
	skippedIds = append(skippedIds, id...)

	for _, us := range userSwipes {
		if us.LikedUserID != nil {
			skippedIds = append(skippedIds, *us.LikedUserID)
		}
	}

	err = r.DB.
		WithContext(ctx).
		Debug().
		Raw(`select u.* from users u 
				where u.deleted_at isnull 
				and u.id not in (?) 
				order by verified_user, id desc
			`, skippedIds).
		Scan(&userModels).
		Error
	if err != nil {
		return []entity.User{}, err
	}
	if len(userModels) < 1 {
		return []entity.User{}, gorm.ErrRecordNotFound
	}

	userEntities := []entity.User{}
	for _, um := range userModels {
		userEntities = append(userEntities, um.ToEntity())
	}

	return userEntities, nil
}
