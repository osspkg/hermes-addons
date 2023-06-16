package hermesaddons

type UserGetter interface {
	ID() uint64
	Alias() string
	Email() string
	Name() string
	Icon() string
}
