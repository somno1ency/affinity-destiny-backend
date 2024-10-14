// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util

func TransformList[S, T any](source []S, transform func(S) T) []T {
	var target []T
	for _, v := range source {
		target = append(target, transform(v))
	}

	return target
}
