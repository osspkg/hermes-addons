/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

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
