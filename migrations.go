package hermesaddons

type MigrationsGetter interface {
	Table() string
	Data() []DatabaseMigration
}

type DatabaseMigration struct {
	ID   string
	Up   string
	Down string
}
