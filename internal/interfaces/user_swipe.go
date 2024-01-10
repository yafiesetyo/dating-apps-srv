package interfaces

import (
	"context"

	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
	"github.com/yafiesetyo/dating-apps-srv/internal/entity"
)

//go:generate mockgen -source=user_swipe.go -destination=../mocks/user_swipe.go

type RCreateUserSwipe interface {
	Do(context.Context, entity.UserSwipe) error
}

type USwipe interface {
	Do(context.Context, uint, uint, constants.SwipeAction) error
}
