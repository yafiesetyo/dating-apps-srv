package utils

import (
	"time"

	"github.com/yafiesetyo/dating-apps-srv/internal/constants"
)

func GetDuration() time.Duration {
	t1 := time.Now()
	t2Str := time.Now().AddDate(0, 0, 1).Format(constants.DoBFormat)
	t2, _ := time.Parse(constants.DoBFormat, t2Str)

	return t2.Sub(t1)
}
