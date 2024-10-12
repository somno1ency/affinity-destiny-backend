// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	cipherStr := h.Sum(nil)

	return hex.EncodeToString(cipherStr)
}

func ValidatePassword(plainPwd, salt, password string) bool {
	return Md5Encode(plainPwd+salt) == password
}

func MakePassword(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}
