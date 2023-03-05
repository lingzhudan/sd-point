package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"sd-point/app/user/service/internal/biz"
	"sd-point/app/user/service/internal/define"
	"strconv"
	"time"
)

type sessionRepo struct {
	data *Data
	log  *log.Helper
}

// NewSessionRepo .
func NewSessionRepo(data *Data, logger log.Logger) biz.SessionRepo {
	return &sessionRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// SetSession 设置登录session
func (r *sessionRepo) SetSession(ctx context.Context, session *biz.Session) (sessionId string, err error) {
	newUuid, err := uuid.NewUUID()
	if err != nil {
		r.log.Errorf("failed to create uuid for session id, error: %v", err)
		return
	}
	sessionId = strconv.Itoa(int(session.UID)) + ":" + newUuid.String()
	sessionStr, err := NewSession(session)
	if err != nil {
		r.log.Errorf("failed to marshal session, error: %v", err)
		return "", err
	}
	if err = r.data.rdb.Set(ctx, sessionId, sessionStr, time.Hour).Err(); err != nil {
		r.log.Errorf("rdb error: %v", err)
	}
	return
}

func NewSession(s *biz.Session) (string, error) {
	js, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

// GetSession 获取登录session
func (r *sessionRepo) GetSession(ctx context.Context, sessionId string) (session *biz.Session, err error) {
	value, err := r.data.rdb.Get(ctx, sessionId).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = define.ErrRecordNotFound
			return
		}
		r.log.Errorf("rdb error: %v", err)
		return
	}
	err = json.Unmarshal([]byte(value), &session)
	if err != nil {
		r.log.Errorf("failed to unmarshal session, error: %v", err)
	}
	return
}

// GetSessionByUID 获取用户所有登录session
func (r *sessionRepo) GetSessionByUID(ctx context.Context, uid uint32) (sessions []*biz.Session, err error) {
	values, err := r.data.rdb.Keys(ctx, strconv.Itoa(int(uid))+":*").Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			err = define.ErrRecordNotFound
			return
		}
		r.log.Errorf("rdb error: %v", err)
		return
	}
	for _, value := range values {
		session := new(biz.Session)
		err = json.Unmarshal([]byte(value), &session)
		if err != nil {
			r.log.Errorf("failed to unmarshal session, error: %v", err)
		}
		sessions = append(sessions, session)
	}
	return
}

// DelSession 删除登录session
func (r *sessionRepo) DelSession(ctx context.Context, sessionId string) (err error) {
	if _, err = r.data.rdb.Del(ctx, sessionId).Result(); err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}
