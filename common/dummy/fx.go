package dummy

import (
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewDummy),
	fx.Provide(ModuleBind),
)

func ModuleBind(
	dummy *DummyImpl,
) Dummy {
	return dummy
}
