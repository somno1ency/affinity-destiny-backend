// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

import (
	"math/rand"
	"strings"
)

func GenCode(count int, isUpperCase bool) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	b := make([]rune, count)
	for i := range count {
		b[i] = letters[rand.Intn(len(letters))]
	}
	result := string(b)
	if isUpperCase {
		result = strings.ToUpper(result)
	}

	return result
}

func GenNumber(count int) string {
	letters := []rune("0123456789")
	b := make([]rune, count)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}
