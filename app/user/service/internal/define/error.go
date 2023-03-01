package define

import (
	"github.com/go-kratos/kratos/v2/errors"
)

var (
	ErrPasswordIncorrect = errors.New(400, "password incorrect, %s", "error")
	ErrRecordNotFound    = errors.New(404, "record not found, %s", "error")
	ErrDuplicateKey      = errors.New(404, "duplicate key, %s", "error")
)
