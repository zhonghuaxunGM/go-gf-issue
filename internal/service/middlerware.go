package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// sMiddleware is service struct of module Middleware.
type sMiddleware struct{}

// insMiddleware is the instance of service Middleware.
var insMiddleware = sMiddleware{}

// Middleware returns the interface of Middleware service.
func Middleware() *sMiddleware {
	return &insMiddleware
}

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	// TODO should replaced
	// customCtx := &model.Context{
	// 	Session: r.Session,
	// }
	// Context().Init(r, customCtx)
	// if user := Session().GetUser(r.Context()); user != nil {
	// 	customCtx.User = &model.ContextUser{
	// 		Id:       user.Id,
	// 		Passport: user.Passport,
	// 		Nickname: user.Nickname,
	// 	}
	// }
	// Continue execution of next middleware.
	r.Middleware.Next()
}

// Auth validates the request to allow only signed-in users visit.
func (s *sMiddleware) Auth(r *ghttp.Request) {
	Auth().MiddlewareFunc()(r)
	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
