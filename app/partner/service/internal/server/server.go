package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/wire"
)

// ProviderSet is user providers.
var ProviderSet = wire.NewSet(NewHTTPServer, NewGRPCServer, NewCronJob)
var (
	ErrorsMsgMap = map[string]string{
		"UNKNOWN_ERROR":        "未知错误",
		"ACCOUNT_EXIST":        "账号已存在",
		"ACCOUNT_ILLEGAL":      "账号只能包含字母数字下划线",
		"USER_REGISTER_FAILED": "用户注册失败",
		"USER_LOGIN_FAILED":    "用户登录失败或账号不存在",
		"USER_DELETE_FAILED":   "用户删除失败",
		"PERMISSION_DENY":      "没有权限",
		"LOGIN_STATE_TIMEOUT":  "登录已过期，请重新登录",
		"USER_LOGOUT_FAILED":   "用户注销失败",
		"USER_TAGS_EMPTY":      "搜索标签不能为空",
		"USER_SEARCH_FAILED":   "用户搜索失败",
		"UPDATE_USER":          "用户信息修改错误",
	}
)

func responseServer() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			reply, err = handler(ctx, req)
			if err != nil {
				e, ok := err.(*errors.Error)
				if ok {
					if m, ok := ErrorsMsgMap[e.Reason]; ok {
						e.Message = m
					}
					return
				}
				e.Message = ErrorsMsgMap["UNKNOWN_ERROR"]
			}
			return
		}
	}
}
