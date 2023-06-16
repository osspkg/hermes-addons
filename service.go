/* 
 *  LGPL-3.0 license
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>
 *  See the full text of the license in the LICENSE file in the root directory.
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
