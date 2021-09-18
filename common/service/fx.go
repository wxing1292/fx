package service

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewService),
	fx.Invoke(lifecycleHooks),
)

func lifecycleHooks(
	lc fx.Lifecycle,
	service *ServiceImpl,
) *ServiceImpl {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			service.logger.Info("logger from service say hi")
			return service.Start()
		},
		OnStop: func(ctx context.Context) error {
			return service.Stop()
		},
	})

	return service
}
