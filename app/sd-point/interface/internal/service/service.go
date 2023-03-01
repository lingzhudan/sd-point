package service

import (
	"context"
	"github.com/google/wire"
	"sd-point/app/sd-point/interface/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewSdPointInterfaceService)

// GetSession 获取session 无则抛出恐慌
func GetSession(ctx context.Context) (s *biz.Session) {
	var ok bool
	if s, ok = ctx.Value("session").(*biz.Session); !ok {
		panic("session type error")
	}
	return
}

// GetSessionID 获取session 无则抛出恐慌
func GetSessionID(ctx context.Context) (s string) {
	var ok bool
	if s, ok = ctx.Value("sessionId").(string); !ok {
		panic("session type error")
	}
	return
}
