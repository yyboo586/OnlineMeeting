package consts

const (
	// CachePrefix 应用缓存数据前缀
	CachePrefix = "APP:"

	CacheModelMem   = "memory"
	CacheModelRedis = "redis"
	CacheModelDist  = "dist"
)

const (
	// PageSize 分页长度
	PageSize = 10
)

type CtxTokenKey string

const TokenKey CtxTokenKey = "token"
