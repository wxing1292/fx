package logger

import "fmt"

type (
	Logger interface {
		Info(msg string)
	}
)

type (
	LoggerImpl struct {
	}
)

func NewLogger() (*LoggerImpl, error) {
	return &LoggerImpl{}, nil
}

func (l *LoggerImpl) Info(msg string) {
	fmt.Printf("## normal logger: %v\n", msg)
}
