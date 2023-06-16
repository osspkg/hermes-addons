/*
 *  LGPL-3.0 license
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>
 *  See the full text of the license in the LICENSE file in the root directory.
 */

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
