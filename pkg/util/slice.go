// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

func Contains[T comparable](slice []T, ele T) bool {
	for _, v := range slice {
		if v == ele {
			return true
		}
	}

	return false
}
