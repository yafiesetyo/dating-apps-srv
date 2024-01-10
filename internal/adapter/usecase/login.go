package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yafiesetyo/dating-apps-srv/config"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
	"github.com/yafiesetyo/dating-apps-srv/internal/interfaces"
	"gorm.io/gorm"
)

type login struct {
	userGetter  interfaces.RFindUserByUsername
	passwordLib interfaces.BCrypt
}

var _ interfaces.ULogin = (*login)(nil)

func NewLogin(
	userGetter interfaces.RFindUserByUsername,
	passwordLib interfaces.BCrypt,
) *login {
	return &login{
		userGetter:  userGetter,
		passwordLib: passwordLib,
	}
}

func (u *login) Do(ctx context.Context, in entity.User) (string, error) {
	user, err := u.userGetter.Do(ctx, in.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.New("username not found")
	}
	if err != nil {
		return "", err
	}

	if err := u.passwordLib.CompareAndHash([]byte(user.Password), []byte(in.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := u.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *login) generateToken(in entity.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":     1,
		"user_id": in.ID,
		"exp":     time.Now().Add(config.Cfg.JWT.TTL).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(config.Cfg.JWT.SecretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
