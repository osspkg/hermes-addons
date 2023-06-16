/*
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
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
