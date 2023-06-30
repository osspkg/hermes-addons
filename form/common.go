/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package form

type draw interface {
	Draw()
}

type elementBlock struct {
	Tag   string   `json:"t"`
	Class []string `json:"c,omitempty"`
	Next  []draw   `json:"n,omitempty"`
}

func (elementBlock) Draw() {}

type elementText struct {
	Tag   string   `json:"t"`
	Class []string `json:"c,omitempty"`
	Value []string `json:"v,omitempty"`
}

func (elementText) Draw() {}

type elementBody struct {
	Get  uint   `json:"get"`
	Set  uint   `json:"set"`
	HTML []draw `json:"html,omitempty"`
}

func (elementBody) Draw() {}
