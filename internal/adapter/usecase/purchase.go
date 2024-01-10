package usecase

import (
	"context"
	"errors"

	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
)

type purchase struct {
	userGetter           interfaces.RFindUserByID
	premiumFeatureSetter interfaces.RSetPremiumFeature
}

var _ interfaces.UPurchase = (*purchase)(nil)

func NewPurchase(
	userGetter interfaces.RFindUserByID,
	premiumFeatureSetter interfaces.RSetPremiumFeature,
) *purchase {
	return &purchase{
		userGetter:           userGetter,
		premiumFeatureSetter: premiumFeatureSetter,
	}
}

func (u *purchase) Do(ctx context.Context, in entity.User) error {
	user, err := u.userGetter.Do(ctx, in.ID)
	if err != nil {
		return err
	}

	if user.UnlimitedSwipe && in.UnlimitedSwipe {
		return errors.New("unlimited swipe already purchased")
	}
	if user.VerifiedUser && in.VerifiedUser {
		return errors.New("verified user already purchased")
	}

	if in.UnlimitedSwipe {
		return u.premiumFeatureSetter.Do(ctx, in.ID, constants.UnlimitedSwipe)
	}

	return u.premiumFeatureSetter.Do(ctx, in.ID, constants.VerifiedUser)
}
