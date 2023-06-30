/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package form_test

import (
	"encoding/json"
	"github.com/osspkg/hermes-addons/form"
	"testing"
)

func TestUnit_Blocks(t *testing.T) {
	nodes := form.Body(1, 2,
		form.Box(
			form.Row(
				form.Col(nil, form.Box()),
				form.Col(&form.ColOpt{Class: []string{"color1"}, Size: 2},
					form.Box(
						form.Text(nil, "Hello", "World"),
					),
				),
			),
		),
	)
	b, err := json.Marshal(nodes)
	if err != nil {
		t.Fail()
	}
	actual := string(b)
	expacted := `{"get":1,"set":2,"html":[{"t":"div","c":["box"],"n":[{"t":"div","c":["row"],"n":[{"t":"div","c":["col"],"n":[{"t":"div","c":["box"]}]},{"t":"div","c":["col-2","color1"],"n":[{"t":"div","c":["box"],"n":[{"t":"p","v":["Hello","World"]}]}]}]}]}]}`
	if expacted != actual {
		t.Log(actual)
		t.Fail()
	}
}
