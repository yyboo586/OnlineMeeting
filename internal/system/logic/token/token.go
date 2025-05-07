package token

import (
	"OnlineMeeting/internal/consts"
	"OnlineMeeting/internal/system/model"
	"OnlineMeeting/internal/system/service"

	"OnlineMeeting/library/liberr"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-token/adapter"
	"github.com/tiger1103/gfast-token/gftoken"
)

type sToken struct {
	*gftoken.GfToken
}

func New() service.IGfToken {
	var (
		ctx = gctx.New()
		opt *model.TokenOptions
		err = g.Cfg().MustGet(ctx, "gfToken").Struct(&opt)
		fun gftoken.OptionFunc
	)
	liberr.ErrIsNil(ctx, err)
	if opt.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis() //redis缓存
	} else if opt.CacheModel == consts.CacheModelDist {
		//磁盘缓存
		fun = gftoken.WithDistConfig(&adapter.Config{
			Dir: opt.DistPath,
		})
	} else {
		fun = gftoken.WithGCache() // 内存缓存
	}
	return &sToken{
		GfToken: gftoken.NewGfToken(
			gftoken.WithCacheKey(opt.CacheKey),
			gftoken.WithTimeout(opt.Timeout),
			gftoken.WithMaxRefresh(opt.MaxRefresh),
			gftoken.WithMultiLogin(opt.MultiLogin),
			gftoken.WithExcludePaths(opt.ExcludePaths),
			fun,
		),
	}
}

func init() {
	service.RegisterGToken(New())
}
