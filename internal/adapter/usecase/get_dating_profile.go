package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	logger "github.com/sirupsen/logrus"
	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"github.com/yafiesetyo/dating-apps-srv/utils"
)

type datingProfileGetter struct {
	cache      interfaces.Cache
	userGetter interfaces.RFindUserByID

	profileGetter interfaces.RFindUserProfile
}

var _ interfaces.UViewDatingProfile = (*datingProfileGetter)(nil)

func NewDatingProfileGetter(
	cache interfaces.Cache,
	userGetter interfaces.RFindUserByID,
	profileGetter interfaces.RFindUserProfile,
) *datingProfileGetter {
	return &datingProfileGetter{
		cache:         cache,
		userGetter:    userGetter,
		profileGetter: profileGetter,
	}
}

func (u *datingProfileGetter) Do(ctx context.Context, in entity.User) (entity.User, error) {
	user, err := u.userGetter.Do(ctx, in.ID)
	if err != nil {
		logger.Errorf("failed to get user, err: %v", err)
		return entity.User{}, err
	}

	if user.UnlimitedSwipe {
		users, err := u.profileGetter.Do(ctx, []uint{in.ID})
		if err != nil {
			return entity.User{}, err
		}

		return users[0], nil
	}

	keys, err := u.cache.Keys(ctx, fmt.Sprintf("view.%d.*", in.ID))
	if err != nil {
		return entity.User{}, err
	}

	if len(keys) > 10 {
		return entity.User{}, errors.New("limit view reached")
	}

	ids := []uint{in.ID}
	for _, key := range keys {
		viewedId, err := strconv.ParseUint(strings.Split(key, ".")[2], 10, 64)
		if err != nil {
			return entity.User{}, err
		}

		ids = append(ids, uint(viewedId))
	}

	datingProfile, err := u.profileGetter.Do(ctx, ids)
	if err != nil {
		return entity.User{}, err
	}

	err = u.cache.Set(ctx, fmt.Sprintf(constants.ViewKey, in.ID, datingProfile[0].ID), true, utils.GetDuration())
	if err != nil {
		return entity.User{}, err
	}

	return datingProfile[0], nil
}
