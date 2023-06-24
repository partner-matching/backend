package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	_ "github.com/go-kratos/swagger-api/openapiv2"
	"github.com/gorilla/handlers"
	v1 "github.com/partner-matching/backend/api/partner/service/v1"
	"github.com/partner-matching/backend/app/partner/service/internal/conf"
	"github.com/partner-matching/backend/app/partner/service/internal/service"
)

// NewHTTPServer new a HTTP user.
func NewHTTPServer(c *conf.Server, userService *service.UserService, partnerService *service.PartnerService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
				l := log.NewHelper(log.With(logger, "message", "panic"))
				l.Error(err)
				return nil
			})),
			ratelimit.Server(),
			responseServer(),
			logging.Server(log.NewFilter(logger, log.FilterLevel(log.LevelInfo))),
			validate.Validator(),
		),
		// 允许跨域
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "DELETE", "PUT", "HEAD"}),
			handlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Content-Type"}),
			handlers.AllowCredentials(),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterUserServiceHTTPServer(srv, userService)
	v1.RegisterPartnerServiceHTTPServer(srv, partnerService)

	// swagger 调试开启
	openAPIHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIHandler)
	return srv
}
