package logger

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewLogger),
	fx.Provide(ModuleBind),
)

func ModuleBind(
	logger *LoggerImpl,
) Logger {
	return logger
}
