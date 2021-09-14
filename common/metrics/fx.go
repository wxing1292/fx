package metrics

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewMetrics),
	fx.Provide(ModuleBind),
)

func ModuleBind(
	metrics *MetricsImpl,
) Metrics {
	return metrics
}
