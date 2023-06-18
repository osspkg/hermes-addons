/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package base

const (
	Icon64  IconSize = 64
	Icon128 IconSize = 128
)

type IconSize uint

type IconGetter interface {
	GetIcon(size IconSize) string
}
