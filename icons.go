/* 
 *  LGPL-3.0 license
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>
 *  See the full text of the license in the LICENSE file in the root directory.
 */

package hermesaddons

const (
	Icon64  IconSize = 64
	Icon128 IconSize = 128
)

type IconSize uint

type IconGetter interface {
	GetIcon(size IconSize) string
}
