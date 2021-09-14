package cache

import (
	"fmt"

	"github.com/wxing1292/fx/common/daemon"
	"github.com/wxing1292/fx/common/logger"
	"github.com/wxing1292/fx/common/metrics"
)

type (
	Cache interface {
		daemon.Daemon
	}
)

type (
	CacheImpl struct {
		metrics metrics.Metrics
		logger  logger.Logger
	}
)

func NewCache(
	metrics metrics.Metrics,
	logger logger.Logger,
) *CacheImpl {
	return &CacheImpl{
		metrics: metrics,
		logger:  logger,
	}
}

func (s *CacheImpl) Start() error {
	fmt.Println("## cache start")
	return nil
}

func (s *CacheImpl) Stop() error {
	fmt.Println("## cache stop")
	return nil
}
