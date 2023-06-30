/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package form

type TextOpt struct {
	Class []string
}

func Text(opt *TextOpt, value ...string) draw {
	el := elementText{
		Tag:   "p",
		Value: value,
	}
	if opt != nil {
		el.Class = append(el.Class, opt.Class...)
	}
	return el
}
