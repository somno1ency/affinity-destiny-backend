// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types

type QueryListReq struct {
	Page       int64  `form:"page,default=1"`
	PageSize   int64  `form:"pageSize,default=20"`
	OrderField string `form:"orderField,default=id"`
	Asc        bool   `form:"asc,default=true"`
}
