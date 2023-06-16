/*
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package hermesaddons

type UserGetter interface {
	ID() uint64
	Alias() string
	Email() string
	Name() string
	Icon() string
}
