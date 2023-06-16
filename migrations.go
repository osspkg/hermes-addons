/* 
 *  LGPL-3.0 license
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>
 *  See the full text of the license in the LICENSE file in the root directory.
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
