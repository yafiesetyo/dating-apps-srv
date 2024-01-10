package usecase

import (
	"context"
	"errors"

	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"github.com/yafiesetyo/dating-apps-srv/internal/logger"
	"gorm.io/gorm"
)

type register struct {
	userChecker interfaces.RFindUserByUsername
	userCreator interfaces.RCreateUser
	passwordLib interfaces.BCrypt
}

var _ interfaces.UCreateUser = (*register)(nil)

func NewRegister(
	userChecker interfaces.RFindUserByUsername,
	userCreator interfaces.RCreateUser,
	passwordLib interfaces.BCrypt,
) *register {
	return &register{
		userChecker: userChecker,
		userCreator: userCreator,
		passwordLib: passwordLib,
	}
}

func (u *register) Do(ctx context.Context, in entity.User) error {
	ctxName := "usecase.register"
	user, err := u.userChecker.Do(ctx, in.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logger.Error(ctxName, "failed to get user: %v", err)
		return err
	}
	if user.ID > 0 {
		return errors.New("username already used")
	}

	pbt, err := u.passwordLib.GenerateFromPassword([]byte(in.Password), 10)
	if err != nil {
		return err
	}
	in.Password = string(pbt)

	if err := u.userCreator.Do(ctx, in); err != nil {
		logger.Error(ctxName, "failed to create user: %v", err)
		return err
	}

	return nil
}
