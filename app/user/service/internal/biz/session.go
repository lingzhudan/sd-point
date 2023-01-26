package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/uuid"
)

type SessionUseCase struct {
	repo SessionRepo
	log  *log.Helper
}

func NewSessionUseCase(repo SessionRepo, logger log.Logger) *SessionUseCase {
	return &SessionUseCase{repo: repo, log: log.NewHelper(logger)}
}

type SessionRepo interface {
	Set(ctx context.Context, key string, value interface{}) (err error)
	Get(ctx context.Context, key string) (value interface{}, err error)
	Del(ctx context.Context, keys []string) (err error)
}

// NewSessionID 生成新的sessionID并使其生效
func (uc *SessionUseCase) NewSessionID(ctx context.Context, user *User) (sessionId string, err error) {
	var newUuid uuid.UUID
	if newUuid, err = uuid.NewUUID(); err != nil {
		uc.log.Errorf("failed to create uuid for session id, error: %v", err)
		return
	}
	sessionId = newUuid.String()
	if err = uc.repo.Set(ctx, sessionId, user.UID); err != nil {
		uc.log.Errorf("failed to set session id to database, error: %v", err)
	}
	return
}

func (uc *SessionUseCase) DelSessionID(ctx context.Context, sessionId string) (err error) {
	if err = uc.repo.Del(ctx, []string{sessionId}); err != nil {
		uc.log.Errorf("failed to delete session id from database, error: %v", err)
	}
	return
}

func (uc *SessionUseCase) Get(ctx context.Context, key string) (value string, err error) {
	var result interface{}
	if result, err = uc.repo.Get(ctx, key); err != nil {
		uc.log.Errorf("failed to get session, error: %v", err)
	}
	value, _ = result.(string)
	return
}

func (uc *SessionUseCase) Set(ctx context.Context, key string, value interface{}) (err error) {
	if err = uc.repo.Set(ctx, key, value); err != nil {
		uc.log.Debugf("failed to set session, error: %v", err)
	}
	return
}
