/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package hermesaddons

import (
	"context"
	"encoding/json"
)

type JsonRPCGetter interface {
	Form(id uint) (json.Marshaler, error)
	Call(ctx context.Context, id uint, data []byte, user UserGetter) (json.Marshaler, error)
}
