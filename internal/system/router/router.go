package router

import (
	"OnlineMeeting/internal/system/controller/login"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Router struct{}

var R = new(Router)

func (r *Router) BindController(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		group.Bind(
			login.NewV1(),
		)
	})
}
