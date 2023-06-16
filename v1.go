package hermesaddons

type AppV1 interface {
	SchemaGetter
	PkgNameGetter
	VersionGetter

	ServiceSetter

	IconGetter
	JsonRPCGetter

	Info() InfoV1
	Database() MigrationsGetter
	ACL() ACLGetter
}

type InfoV1 struct {
	Name        string
	Author      string
	Email       string
	Title       string
	Description string
	Info        string
}
