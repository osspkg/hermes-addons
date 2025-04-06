/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package addons

import (
	"fmt"
	"strings"

	"go.osspkg.com/do"
	"go.osspkg.com/goppy/v2/orm"

	"go.osspkg.com/hermes-addons/internal"
)

func (a *_addons) WithDatabase(tags []string) {
	tags = do.Unique[string](tags)
	envs := make([]Env, 0, 10)
	conf := `
pgsql:
`
	for _, tag := range tags {
		envSuffix := internal.BuildEnv(tag)

		conf += fmt.Sprintf(`
    - tags: %[1]s
      host: @env(DB_%[2]s_HOST#127.0.0.1)
      port: @env(DB_%[2]s_PORT#5432)
      user: @env(DB_%[2]s_USER#postgres)
      password: @env(DB_%[2]s_PASSWRD#postgres)
      schema: @env(DB_%[2]s_SCHEMA#postgres)
      app_name: %[3]s
`, tag, envSuffix, strings.TrimSpace(a.info.Name))

		envs = append(envs,
			Env{Key: "DB_" + envSuffix + "_HOST", Description: "PostgresSQL HOST for tag: " + tag, Default: "127.0.0.1"},
			Env{Key: "DB_" + envSuffix + "_PORT", Description: "PostgresSQL PORT for tag: " + tag, Default: "5432"},
			Env{Key: "DB_" + envSuffix + "_USER", Description: "PostgresSQL USER for tag: " + tag, Default: "postgres"},
			Env{Key: "DB_" + envSuffix + "_PASSWRD", Description: "PostgresSQL PASSWORD for tag: " + tag, Default: "postgres"},
			Env{Key: "DB_" + envSuffix + "_SCHEMA", Description: "PostgresSQL SCHEMA for tag: " + tag, Default: internal.BuildSchema(a.info.Package)},
		)
	}

	a.WithDependencies(Dep{
		Inject: []any{
			orm.WithORM(),
			orm.WithPgsqlClient(),
		},
		Config: conf,
		Envs:   envs,
	})
}
