package define

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrRecordNotFound = errors.New(404, "record not found, %s", "error")
	ErrDuplicateKey   = errors.New(500, "duplicate key, %s", "error")

	ErrUserNotFound          = errors.New(404, "user not found, %s", "error")
	ErrPasswordIncorrect     = errors.New(400, "password incorrect, %s", "error")
	ErrAccountRegistered     = errors.New(500, "account registered, %s", "error")
	ErrWechatRegistered      = errors.New(500, "wechat registered, %s", "error")
	ErrPhoneNumberRegistered = errors.New(500, "phone number registered, %s", "error")
)
