/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package addons

import (
	"go.osspkg.com/console"
	"go.osspkg.com/ioutils/data"
)

type _addons struct {
	info Manifest
	envs Envs
	deps []any
	conf *data.Buffer
	cli  *console.Console
}

func New(m Manifest) Addons {
	a := &_addons{
		cli:  console.New("hermes.addon", "hermes addon"),
		conf: data.NewBuffer(1024),
		info: m,
	}

	configBase(a.conf)

	return a
}

func (a *_addons) Run() {
	a.cli.RootCommand(a.commandApp())
	a.cli.AddCommand(a.commandSetup())
	a.cli.Exec()
}

func (a *_addons) WithDependencies(deps ...Dep) {
	for _, dep := range deps {
		a.deps = append(a.deps, dep.Inject...)
		a.envs = append(a.envs, dep.Envs...)
		confWrite(a.conf, dep.Config)
	}
}
