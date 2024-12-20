// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel

		FindOneByCustomId(ctx context.Context, customId string) (*User, error)
		FindByUserIds(ctx context.Context, userIds []int64) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserModel) FindOneByCustomId(ctx context.Context, customId string) (*User, error) {
	var resp User
	query := fmt.Sprintf("select %s from %s where `CustomId` = ? limit 1", userRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, customId)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserModel) FindByUserIds(ctx context.Context, userIds []int64) ([]*User, error) {
	sql := "select %s from %s where `Id` in ("
	placeHolder := ""
	for i, id := range userIds {
		if i > 0 && (i != len(userIds)) {
			placeHolder += ", "
		}
		placeHolder += fmt.Sprintf("%d", id)
	}
	sql += placeHolder + ")"
	query := fmt.Sprintf(sql, userRows, m.table)
	var resp []*User
	err := m.conn.QueryRowsCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
