/*
 *  Copyright (c) 2023-2025 Mikhail Knyazhev <markus621@yandex.com>. All rights reserved.
 *  Use of this source code is governed by a BSD 3-Clause license that can be found in the LICENSE file.
 */

package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"regexp"
	"strings"
)

var rex = regexp.MustCompile(`(?mUi)[^0-9a-z\_]+`)

func BuildEnv(s string) string {
	s = strings.TrimSpace(s)
	s = rex.ReplaceAllString(s, "")
	return strings.ToUpper(s)
}

func BuildSchema(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	code := hex.EncodeToString(h.Sum(nil))
	s = strings.TrimSpace("addon_" + code)
	s = rex.ReplaceAllString(s, "")
	return strings.ToLower(s)
}
