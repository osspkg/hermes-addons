/*
 *  Copyright (c) 2023 Mikhail Knyazhev <markus621@yandex.ru>. All rights reserved.
 *  Use of this source code is governed by a LGPL-3.0 license that can be found in the LICENSE file.
 */

package hermesaddons

type AppV1 interface {
	SchemaGetter
	PkgNameGetter
	VersionGetter

	ServiceSetter

	IconGetter
	JsonRPCGetter

	Info() InfoV1
	Database() MigrationsGetter
	ACL() ACLGetter
}

type InfoV1 struct {
	Name        string
	Author      string
	Email       string
	Title       string
	Description string
	Info        string
}
