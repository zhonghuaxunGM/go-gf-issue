package cmd

import (
	"context"
	"go-gf-issue/internal/controller"
	"go-gf-issue/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetPort(7799)
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Group("/v1", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						controller.UserInfo,
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
