package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
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
				group.ALLMap(g.Map{
					"/hello":  hello,
					"/hello2": hello,
				})
			})
			s.BindHandler("/ws", func(r *ghttp.Request) {
				// var ctx = r.Context()
				ws, err := r.WebSocket()
				if err != nil {
					glog.Error(context.Background(), err)
					r.Exit()
				}
				ctx, cancelFunc := context.WithCancel(context.Background())
				r.SetCtx(ctx)
				ws.SetCloseHandler(func(code int, text string) error {
					fmt.Println("code, text:", code, text)
					cancelFunc()
					return ws.Close()
				})
				var times int
				for {
					fmt.Println("repeatedly times:", times)
					select {
					case <-r.GetCtx().Done():
						fmt.Println("done")
						err = ws.Close()
						if err != nil {
							fmt.Println("ws.Close()", err)
						}
						return
					default:
						fmt.Println("get data from text")
						<-time.After(time.Second)
					}
					times++
				}
			})
			s.Run()
			return nil
		},
	}
)

func hello(r *ghttp.Request) {}
