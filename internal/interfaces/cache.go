package interfaces

import (
	"context"
	"time"
)

//go:generate mockgen -source=cache.go -destination=../mocks/cache.go

type Cache interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, interface{}, time.Duration) error
	Keys(context.Context, string) ([]string, error)
}
