package interfaces

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
)

//go:generate mockgen -source=user.go -destination=../mocks/user.go

// Repository
type RCreateUser interface {
	Do(context.Context, entity.User) error
}

type RFindUserByID interface {
	Do(context.Context, uint) (entity.User, error)
}

type RFindUserProfile interface {
	Do(context.Context, []uint) ([]entity.User, error)
}

type RFindUserByUsername interface {
	Do(context.Context, string) (entity.User, error)
}

type RSetPremiumFeature interface {
	Do(context.Context, uint, constants.Features) error
}

// Usecase
type ULogin interface {
	Do(context.Context, entity.User) (string, error)
}

type UCreateUser interface {
	Do(context.Context, entity.User) error
}

type UViewProfile interface {
	Do(context.Context, entity.User) (entity.User, error)
}

type UViewDatingProfile interface {
	Do(context.Context, entity.User) (entity.User, error)
}

type UPurchase interface {
	Do(context.Context, entity.User) error
}
