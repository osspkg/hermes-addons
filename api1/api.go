/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package api1

import "github.com/osspkg/hermes-addons/base"

type Api interface {
	base.SchemaGetter
	base.PkgNameGetter
	base.VersionGetter
	base.ServiceSetter
	base.IconGetter
	Database() base.MigrationsGetter

	JsonRPCGetter
	Info() Info
	ACL() ACLGetter
}
