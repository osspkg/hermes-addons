/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
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
