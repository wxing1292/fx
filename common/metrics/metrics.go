package metrics

import (
	"github.com/wxing1292/fx/common/dummy"
	"github.com/wxing1292/fx/common/logger"
)

type (
	Metrics interface {
		IncCount(int)
	}
)

type (
	MetricsImpl struct {
		Logger logger.Logger
		dummy  dummy.Dummy
	}
)

func NewMetrics(
	logger logger.Logger,
	dummy dummy.Dummy,
) (*MetricsImpl, error) {
	return &MetricsImpl{
		Logger: logger,
		dummy:  dummy,
	}, nil
}

func (m *MetricsImpl) IncCount(int) {}
