/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

import (
	"go.osspkg.com/ioutils/pool"
)

var Pool = pool.New[*Transport](NewTransport)
