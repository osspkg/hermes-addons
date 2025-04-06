/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package errs

import "errors"

var (
	ErrNotFound         = errors.New("not found")
	ErrAlreadyExist     = errors.New("already exist")
	ErrMethodNotAllowed = errors.New("method not allowed")
	ErrInvalidRequest   = errors.New("invalid request")
)
