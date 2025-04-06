/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package addons

//go:generate easyjson

type Addons interface {
	WithDependencies(deps ...Dep)
	WithDatabase(tags []string)
	Run()
}

type Type uint64

const (
	TypeApp Type = 0
)

//easyjson:json
type Manifest struct {
	Name        string `json:"name"`
	Package     string `json:"package"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Version     string `json:"version"`
	Type        Type   `json:"type"`
	Links       []Link `json:"links"`
	Menu        Menu   `json:"menu"`
}

//easyjson:json
type Link struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

//easyjson:json
type Env struct {
	Key         string `json:"k"`
	Description string `json:"d"`
	Default     string `json:"v"`
}

//easyjson:json
type Envs []Env

//easyjson:json
type Menu struct {
	Group string `json:"group"`
	Title string `json:"title"`
}

type Dep struct {
	Inject []any
	Config string
	Envs   []Env
}
