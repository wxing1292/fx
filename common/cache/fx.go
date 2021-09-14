package cache

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewCache),
	fx.Provide(ModuleBind),
	fx.Invoke(lifecycleHooks),
)

func ModuleBind(
	logger *CacheImpl,
) Cache {
	return logger
}

// this completely decouples
// cache with entiry which depends on cache
// e.g. entity depending on cache does not need to call start / stop on cache
// but rather just use cache
func lifecycleHooks(
	lc fx.Lifecycle,
	cache Cache,
) Cache {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return cache.Start()
		},
		OnStop: func(ctx context.Context) error {
			return cache.Stop()
		},
	})

	return cache
}
