package daemon

type (
	Daemon interface {
		Start() error
		Stop() error
	}
)
