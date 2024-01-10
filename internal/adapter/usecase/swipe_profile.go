package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"github.com/yafiesetyo/dating-apps-srv/utils"
	"gorm.io/gorm"
)

type swipeProfile struct {
	cache             interfaces.Cache
	userGetter        interfaces.RFindUserByID
	unlikedUserGetter interfaces.RFindUserProfile
	swipeCreator      interfaces.RCreateUserSwipe
}

var _ interfaces.USwipe = (*swipeProfile)(nil)

func NewSwipeProfile(
	cache interfaces.Cache,
	userGetter interfaces.RFindUserByID,
	swipeCreator interfaces.RCreateUserSwipe,
	unlikedUserGetter interfaces.RFindUserProfile,
) *swipeProfile {
	return &swipeProfile{
		cache:             cache,
		userGetter:        userGetter,
		swipeCreator:      swipeCreator,
		unlikedUserGetter: unlikedUserGetter,
	}
}

func (u *swipeProfile) Do(ctx context.Context, id, swipe uint, action constants.SwipeAction) error {
	if id == swipe {
		return errors.New("cannot swipe yourself")
	}

	user, err := u.userGetter.Do(ctx, id)
	if err != nil {
		return err
	}

	likedUser, err := u.userGetter.Do(ctx, swipe)
	if err != nil {
		return err
	}
	if likedUser.ID < 1 {
		return gorm.ErrRecordNotFound
	}

	unlikedUsers, err := u.unlikedUserGetter.Do(ctx, []uint{id})
	if err != nil {
		return err
	}

	isAvailable := false
	for _, lu := range unlikedUsers {
		if lu.ID == swipe {
			isAvailable = true
			break
		}
	}
	if !isAvailable {
		return errors.New("already swiped")
	}

	if user.UnlimitedSwipe {
		if action == constants.Like {
			return u.swipeCreator.Do(ctx, entity.UserSwipe{
				UserID:      id,
				LikedUserID: &swipe,
			})
		}

		return nil
	}

	res, err := u.cache.Get(ctx, fmt.Sprintf(constants.SwipeCounterKey, id))
	if err != nil {
		return err
	}
	if res == "" {
		res = "0"
	}

	counter, err := strconv.ParseUint(res, 10, 64)
	if err != nil {
		return err
	}

	if counter == 10 {
		return errors.New("swipe limit reached")
	}
	counter++

	if action == constants.Like {
		err = u.swipeCreator.Do(ctx, entity.UserSwipe{
			UserID:      id,
			LikedUserID: &swipe,
		})
		if err != nil {
			return err
		}
	}

	err = u.cache.Set(ctx, fmt.Sprintf(constants.SwipeCounterKey, id), counter, utils.GetDuration())
	if err != nil {
		return err
	}

	return nil
}
