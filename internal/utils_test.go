/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

import "testing"

func TestUnit_BuildSchema(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Case1",
			arg:  "app.go_do",
			want: "addon_bec61b22fc93923fdbb7b288fbae4f6ce06f78f9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildSchema(tt.arg)
			if got != tt.want {
				t.Errorf("BuildSchema() = %v, want %v", got, tt.want)
			}
			t.Logf("%s len %d", got, len(got))
			if len(got) > 63 {
				t.Errorf("BuildSchema() len %v, want %v", len(got), 63)
			}
		})
	}
}

func TestBuildEnv(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{
			name: "Case1",
			arg:  "app.go_do",
			want: "APPGO_DO",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildEnv(tt.arg); got != tt.want {
				t.Errorf("BuildEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
