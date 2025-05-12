package cmd

import (
	"OnlineMeeting/internal/router"
	"context"

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
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.R.BindController(ctx, group)
			})

			s.SetClientMaxBodySize(100 * 1024 * 1024) // 600MB max file size

			s.Run()
			return nil
		},
	}
)
