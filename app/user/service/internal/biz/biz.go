package biz

import (
	"github.com/google/wire"
	"strings"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUserUseCase, NewSessionUseCase)

type UserCond struct {
	Begin        int
	Count        int
	UIDs         []uint32
	Accounts     []string
	Passwords    []string
	OpenIDs      []string
	PhoneNumbers []string
}

func (c *UserCond) ParseCond() (whereStage string, args []interface{}) {
	var whereStages []string
	if len(c.UIDs) != 0 {
		whereStages = append(whereStages, "`uid` IN ?")
		args = append(args, c.UIDs)
	}
	if len(c.Accounts) != 0 {
		whereStages = append(whereStages, "`account` IN ?")
		args = append(args, c.Accounts)
	}
	if len(c.Passwords) != 0 {
		whereStages = append(whereStages, "`password` IN ?")
		args = append(args, c.Accounts)
	}
	if len(c.OpenIDs) != 0 {
		whereStages = append(whereStages, "`open_id` IN ?")
		args = append(args, c.OpenIDs)
	}
	if len(c.PhoneNumbers) != 0 {
		whereStages = append(whereStages, "`phone_number` IN ?")
		args = append(args, c.PhoneNumbers)
	}
	return strings.Join(whereStages, " AND "), args
}
