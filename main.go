package main

import (
	"github.com/wxing1292/fx/common/cache"
	"github.com/wxing1292/fx/common/logger"
	"github.com/wxing1292/fx/common/metrics"
	"github.com/wxing1292/fx/common/service"
	"go.uber.org/fx"
)

var ServiceModule = fx.Options(
	logger.Module,
	metrics.Module,
	cache.Module,
	service.Module,
)

func main() {
	app := fx.New(ServiceModule)

	// Run starts the application, blocks on the signals channel, and then
	// gracefully shuts the application down.
	app.Run()
}
