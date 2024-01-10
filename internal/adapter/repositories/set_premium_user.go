package repositories

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/adapter/repositories/model"
	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"gorm.io/gorm"
)

type setPremiumUser repo

var _ interfaces.RSetPremiumFeature = (*setPremiumUser)(nil)

func NewSetPremiumUser(
	db *gorm.DB,
) *setPremiumUser {
	return &setPremiumUser{
		DB: db,
	}
}

func (r *setPremiumUser) Do(ctx context.Context, id uint, in constants.Features) error {
	var (
		user model.User
	)

	exec := r.DB.Table(user.TableName()).Where(`"id"=?`, id)
	if in == constants.VerifiedUser {
		return exec.UpdateColumn("verified_user", true).Error
	}

	return exec.UpdateColumn("unlimited_swipe", true).Error
}
