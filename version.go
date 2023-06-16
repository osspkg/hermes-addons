package hermesaddons

import (
	"golang.org/x/mod/semver"
)

type SchemaGetter interface {
	Schema() uint
}

type VersionGetter interface {
	Version() SemVersion
}

type SemVersion string

func (v SemVersion) IsValid() bool {
	return semver.IsValid(string(v))
}

func (v SemVersion) Compare(n string) int {
	return semver.Compare(string(v), n)
}
