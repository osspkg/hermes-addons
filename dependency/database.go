package dependency

import "context"

type (
	ORM interface {
		ExecContext(ctx context.Context, call func(q Executor)) error
		QueryContext(ctx context.Context, call func(q Querier)) error
		TransactionContext(ctx context.Context, call func(v Tx)) error
	}

	Result interface {
		RowsAffected() int64
		LastInsertId() int64
	}

	Executor interface {
		SQL(query string, args ...interface{})
		Params(args ...interface{})
		Bind(call func(result Result) error)
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
