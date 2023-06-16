/*
 *  LGPL-3.0 license
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>
 *  See the full text of the license in the LICENSE file in the root directory.
 */

package hermesaddons

type ACLGetter interface {
	Setup() []ACLModel
}

type ACLModel struct {
	ID      uint
	Title   string
	FormIDs []uint
}
