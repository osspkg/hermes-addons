/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package main

import (
	"context"
	"net/http"

	"go.osspkg.com/goppy/v2/orm"

	addons "go.osspkg.com/hermes-addons"
	"go.osspkg.com/hermes-addons/rpc"
)

func main() {

	app := addons.New(addons.Manifest{
		Name:        "simple",
		Package:     "com.example.simple",
		Description: "Simple Addon",
		Author:      "UserName <user@example.com>",
		Version:     "v0.0.0-dev",
		Type:        addons.TypeApp,
		Links: []addons.Link{
			{Url: "https://example.com", Description: "Web"},
			{Url: "https://tg/example", Description: "Telegram"},
		},
		Menu: addons.Menu{
			Group: "Application",
			Title: "Simple App",
		},
	})

	app.WithDatabase([]string{"master", "slave"})

	app.WithDependencies(addons.Dep{Inject: []any{NewController}})
	app.Run()
}

type Controller struct {
	orm orm.ORM
}

func NewController(db orm.ORM, srv rpc.Server) error {
	c := &Controller{orm: db}
	srv.AddHandler("com.example.simple.hello", c.Hello)
	return nil
}

func (c *Controller) Hello(_ context.Context, w rpc.Writer, r rpc.Reader) error {
	w.SetCode(http.StatusOK)
	return w.SetField(rpc.FieldBody, []byte("Hello"))
}
