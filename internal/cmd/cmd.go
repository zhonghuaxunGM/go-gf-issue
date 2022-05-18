package cmd

import (
	"context"
	"go-gf-issue/internal/controller"
	"go-gf-issue/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/protocol/goai"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of authz",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// var cfg = gctx.New()
			// err = initCfg(cfg)
			// if err != nil {
			// 	g.Dump(err)
			// 	return
			// }
			// initDBEnv()
			// initSSOEnv()
			// err = initDB()
			// if err != nil {
			// 	g.Dump(err)
			// 	fmt.Printf("DB GetConnection err: %s \n", err.Error())
			// 	return
			// }
			// err = initCasbinDB()
			// if err != nil {
			// 	g.Dump(err)
			// 	fmt.Printf("casbinDB GetEnforcer err: %s \n", err.Error())
			// 	return
			// }
			s := g.Server()
			s.SetPort(7799)
			// TODO 增加路由请求的信息 https://goframe.org/pages/viewpage.action?pageId=17207121
			s.Group("/", func(group *ghttp.RouterGroup) {
				// Group middlewares.
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
					ghttp.MiddlewareHandlerResponse,
				)
				// self-authn
				// group.Group("/authn", func(group *ghttp.RouterGroup) {
				// 	group.Bind(
				// 		controller.Auth,
				// 	)
				// 	// sso connector list
				// 	group.Group("/sso", func(group *ghttp.RouterGroup) {
				// 		group.GET("/connector/list", controller.Cli.ListConnectors)
				// 		group.ALL("/callback", controller.Cli.Callback)
				// 		// group.GET("/ldap/login", controller.Cli.LdapLogin)
				// 		group.GET("/product/link/list", controller.Cli.ListProductInfo)
				// 	})
				// })

				// Register route handlers.
				group.Group("/authz", func(group *ghttp.RouterGroup) {
					// group.Group("/system", func(group *ghttp.RouterGroup) {
					// 	group.Bind(
					// 		controller.SystemInfo,
					// 	)
					// })
					// Special handler that needs authentication.
					group.Group("/v1", func(group *ghttp.RouterGroup) {
						group.Middleware(service.Middleware().Auth)
						group.Bind(
							controller.UserInfo,
							// controller.GroupInfo,
							// controller.RoleInfo,
							// controller.FieldInfo,
							// controller.CategoryInfo,
							// controller.OrgInfo,
						)
						// group.GET("/user/resource", controller.UserInfo.Resource)
					})
				})
				// group.Middleware(service.Middleware().Auth)
				// group.ALL("/auth", controller.CasbinInstanceInfo.Judge)
			})
			// Custom enhance API document.
			enhanceOpenAPIDoc(s)
			// err = controller.InitSSO(s)
			// if err != nil {
			// 	g.Dump(err)
			// 	fmt.Printf("InitSSO err: %s \n", err.Error())
			// }
			// Just run the server.
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       "consts.OpenAPITitle",
		Description: "consts.OpenAPIDescription",
		Contact: &goai.Contact{
			Name: "AuthZ",
			URL:  "http://192.168.3.97:3000",
		},
	}

	// Sort the tags in custom sequence.
	openapi.Tags = &goai.Tags{
		{Name: "consts.OpenAPITagNameAuthZ"},
	}
}
