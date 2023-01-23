package biz

import (
	"github.com/google/wire"
	"strings"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase)

type UserCond struct {
	Begin int
	Count int
	UIDs  []int32
}

func (c *UserCond) ParseCond() (whereStage string, args []interface{}) {
	var builder strings.Builder
	if len(c.UIDs) != 0 {
		builder.WriteString("`uid` IN (?)")
		args = append(args, c.UIDs)
	}
	return builder.String(), args
}
