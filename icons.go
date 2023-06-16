package hermesaddons

const (
	Icon64  IconSize = 64
	Icon128 IconSize = 128
)

type IconSize uint

type IconGetter interface {
	GetIcon(size IconSize) string
}
