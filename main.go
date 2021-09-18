package main

import (
	"github.com/wxing1292/fx/common/cache"
	"github.com/wxing1292/fx/common/logger"
	"github.com/wxing1292/fx/common/service"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	SpecialMetrics, // metrics.Module,
	logger.Module,
	cache.Module,
	service.Module,
	fx.NopLogger, // disable FX's own noizy logger
)

func main() {
	app := fx.New(ServiceModule)

	// Run starts the application, blocks on the signals channel, and then
	// gracefully shuts the application down.
	app.Run()
}
