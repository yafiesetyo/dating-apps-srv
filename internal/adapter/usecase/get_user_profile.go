package usecase

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
)

type userProfileGetter struct {
	getter interfaces.RFindUserByID
}

var _ interfaces.UViewProfile = (*userProfileGetter)(nil)

func NewUserProfileGetter(
	getter interfaces.RFindUserByID,
) *userProfileGetter {
	return &userProfileGetter{getter: getter}
}

func (u *userProfileGetter) Do(ctx context.Context, in entity.User) (entity.User, error) {
	return u.getter.Do(ctx, in.ID)
}
