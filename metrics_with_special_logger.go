package main

import (
	"context"
	"fmt"

	"github.com/wxing1292/fx/common/logger"
	"github.com/wxing1292/fx/common/metrics"
	"go.uber.org/fx"
)

type SpecialLogger struct {
}

func (l *SpecialLogger) Info(msg string) {
	fmt.Printf("## special logger: %v\n", msg)
}

func GetSpecialLogger() logger.Logger {
	return &SpecialLogger{}
}

func UseMetricsLogger(
	lc fx.Lifecycle,
	metric metrics.Metrics,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			metric.(*metrics.MetricsImpl).Logger.Info("logger from metrics say hi")
			return nil
		},
	})
}

var SpecialMetrics = fx.Options(
	// creat & provide the annotated logger
	fx.Provide(
		fx.Annotated{Name: "metrics-logger", Target: GetSpecialLogger},
	),
	// annotate metrics constructor with above annotation
	fx.Provide(fx.Annotate(metrics.NewMetrics, fx.ParamTags(
		`name:"metrics-logger"`,
	))),
	// remaining binding for metrics
	fx.Provide(metrics.ModuleBind),
	// invoke metrics and use logger within
	fx.Invoke(UseMetricsLogger),
)
