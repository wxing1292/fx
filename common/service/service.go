package service

import (
	"fmt"

	"github.com/wxing1292/fx/common/cache"
	"github.com/wxing1292/fx/common/logger"
	"github.com/wxing1292/fx/common/metrics"
)

type (
	ServiceImpl struct {
		cache   cache.Cache
		metrics metrics.Metrics
		logger  logger.Logger
	}
)

func NewService(
	cache cache.Cache,
	metrics metrics.Metrics,
	logger logger.Logger,
) *ServiceImpl {
	return &ServiceImpl{
		cache:   cache,
		metrics: metrics,
		logger:  logger,
	}
}

func (s *ServiceImpl) Start() error {
	fmt.Println("## service start")
	return nil
}

func (s *ServiceImpl) Stop() error {
	fmt.Println("## service stop")
	return nil
}
