package define

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrAccountNotFound   = errors.New(404, "account not found", "")
	ErrUserNotFound      = errors.New(404, "user not found", "")
	ErrSessionNotFound   = errors.New(404, "session not found", "")
	ErrPasswordError     = errors.New(403, "password error", "")
	ErrAccountRegistered = errors.New(500, "account registered", "")
	ErrWechatRegistered  = errors.New(500, "wechat registered", "")
)
