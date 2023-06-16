/* 
 *  LGPL-3.0 license
 *  Copyright (c) 2023 Mikhail Knyzhev <markus621@yandex.ru>
 *  See the full text of the license in the LICENSE file in the root directory.
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
