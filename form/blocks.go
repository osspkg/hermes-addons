/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package form

import "fmt"

func Box(next ...draw) draw {
	return elementBlock{
		Tag:   "div",
		Class: []string{"box"},
		Next:  next,
	}
}

func Row(next ...draw) draw {
	return elementBlock{
		Tag:   "div",
		Class: []string{"row"},
		Next:  next,
	}
}

type ColOpt struct {
	Class []string
	Size  uint
}

func Col(opt *ColOpt, next ...draw) draw {
	el := elementBlock{
		Tag:   "div",
		Class: []string{"col"},
		Next:  next,
	}
	if opt != nil {
		if opt.Size > 0 {
			el.Class = []string{fmt.Sprintf("col-%d", opt.Size)}
		}
		el.Class = append(el.Class, opt.Class...)
	}
	return el
}
