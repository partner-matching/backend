package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/wire"
	"strings"
)

// ProviderSet is user providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewCronJob)
var (
	ErrorsMsgMap = map[string]string{
		"UNKNOWN_ERROR": "未知错误",
	}
)

func responseServer() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if err != nil {
				e, ok := err.(*errors.Error)
				if ok {
					e.Message = strings.SplitN(e.Message, ":", 2)[0]
					return
				}
				e.Message = ErrorsMsgMap["UNKNOWN_ERROR"]
			}
			return
		}
	}
}
