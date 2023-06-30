/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package dependency

import "context"

type (
	ORM interface {
		ExecContext(ctx context.Context, call func(q Executor)) error
		QueryContext(ctx context.Context, call func(q Querier)) error
		TransactionContext(ctx context.Context, call func(v Tx)) error
	}

	Executor interface {
		SQL(query string, args ...interface{})
		Params(args ...interface{})
		Bind(call func(rowsAffected, lastInsertId int64) error)
	}

	Scanner interface {
		Scan(args ...interface{}) error
	}

	Querier interface {
		SQL(query string, args ...interface{})
		Bind(call func(bind Scanner) error)
	}

	Tx interface {
		Exec(vv ...func(e Executor))
		Query(vv ...func(q Querier))
	}
)
