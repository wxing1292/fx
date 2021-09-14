package service

import (
	"fmt"

	"github.com/wxing1292/fx/common/logger"
	"github.com/wxing1292/fx/common/metrics"
)

type (
	ServiceImpl struct {
		metrics metrics.Metrics
		logger  logger.Logger
	}
)

func NewService(
	metrics metrics.Metrics,
	logger logger.Logger,
) *ServiceImpl {
	return &ServiceImpl{
		metrics: metrics,
		logger:  logger,
	}
}

func (s *ServiceImpl) Start() error {
	fmt.Println("service start")
	return nil
}

func (s *ServiceImpl) Stop() error {
	fmt.Println("service stop")
	return nil
}
