// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group_contact

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ GroupContactModel = (*customGroupContactModel)(nil)

type (
	// GroupContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupContactModel.
	GroupContactModel interface {
		groupContactModel
		withSession(session sqlx.Session) GroupContactModel
	}

	customGroupContactModel struct {
		*defaultGroupContactModel
	}
)

// NewGroupContactModel returns a model for the database table.
func NewGroupContactModel(conn sqlx.SqlConn) GroupContactModel {
	return &customGroupContactModel{
		defaultGroupContactModel: newGroupContactModel(conn),
	}
}

func (m *customGroupContactModel) withSession(session sqlx.Session) GroupContactModel {
	return NewGroupContactModel(sqlx.NewSqlConnFromSession(session))
}
