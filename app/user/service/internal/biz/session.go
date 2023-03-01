package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type SessionUseCase struct {
	repo SessionRepo
	log  *log.Helper
}

func NewSessionUseCase(repo SessionRepo, logger log.Logger) *SessionUseCase {
	return &SessionUseCase{repo: repo, log: log.NewHelper(logger)}
}

type SessionRepo interface {
	SetSession(ctx context.Context, session *Session) (sessionId string, err error)
	GetSession(ctx context.Context, sessionId string) (session *Session, err error)
	GetSessionByUID(ctx context.Context, uid uint32) (sessions []*Session, err error)
	DelSession(ctx context.Context, sessionId string) (err error)
}

type Session struct {
	// 用户编号
	UID uint32
}

// NewSessionID 生成新的sessionID并使其生效
func (uc *SessionUseCase) NewSessionID(ctx context.Context, uid uint32) (sessionId string, err error) {
	return uc.repo.SetSession(ctx, &Session{UID: uid})
}

// GetSession 获取sessionID对应的数据
func (uc *SessionUseCase) GetSession(ctx context.Context, sessionId string) (session *Session, err error) {
	return uc.repo.GetSession(ctx, sessionId)
}

// DelSessionID 删除登录session
func (uc *SessionUseCase) DelSessionID(ctx context.Context, sessionId string) (err error) {
	return uc.repo.DelSession(ctx, sessionId)
}
