/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package hermesaddons

import "context"

type ServiceSetter interface {
	Inject(dic DIContainer) error
	Dependency() []string
	Up(ctx context.Context) error
	Down() error
}

type DIContainer interface {
	Get(v string) (interface{}, error)
}
