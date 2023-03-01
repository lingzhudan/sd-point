package middleware

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	pb "sd-point/api/sd-point/interface/v1"
	"sd-point/app/sd-point/interface/internal/biz"
)

func CheckAuthorization(uc *biz.UserUseCase) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				// Do something on entering
				var sessionId string
				var session *biz.Session
				if sessionId = tr.RequestHeader().Get("Authorization"); len(sessionId) == 0 {
					return nil, pb.ErrorNotLoggedIn("not logged in")
				}
				if session, err = uc.GetSession(ctx, sessionId); err != nil {
					return nil, err
				}
				ctx = context.WithValue(ctx, "session", session)
				ctx = context.WithValue(ctx, "sessionId", sessionId)
				defer func() {
					// Do something on exiting
				}()
			}
			return handler(ctx, req)
		}
	}
}
