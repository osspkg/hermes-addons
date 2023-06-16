package hermesaddons

import "context"

type ServiceSetter interface {
	Inject(dic DIContainer) error
	Dependency() []string
	Up(ctx context.Context) error
	Down() error
}

type DIContainer interface {
	Get(v string) (interface{}, error)
}
