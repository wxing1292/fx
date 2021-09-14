package metrics

import "github.com/wxing1292/fx/common/logger"

type (
	Metrics interface {
		IncCount(int)
	}
)

type (
	MetricsImpl struct {
		logger logger.Logger
	}
)

func NewMetrics(
	logger logger.Logger,
) (*MetricsImpl, error) {
	return &MetricsImpl{
		logger: logger,
	}, nil
}

func (m *MetricsImpl) IncCount(int) {}
