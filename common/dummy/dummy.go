package dummy

import "github.com/wxing1292/fx/common/logger"

type (
	Dummy interface{}
)

type (
	DummyImpl struct {
		Logger logger.Logger
	}
)

func NewDummy() (*DummyImpl, error) {
	return &DummyImpl{}, nil
}
