// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user_contact

import (
	"context"
	"fmt"

	"ad.com/pkg/exception"
	"ad.com/pkg/shared/types"
	"ad.com/pkg/util"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserContactModel = (*customUserContactModel)(nil)

type (
	// UserContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserContactModel.
	UserContactModel interface {
		userContactModel
		withSession(session sqlx.Session) UserContactModel

		Find(ctx context.Context, ownerId int64, req *types.QueryListReq) ([]*UserContact, error)
		Count(ctx context.Context, ownerId int64) (int64, error)
		FindDstContact(ctx context.Context, ownerId int64, dstId int64) (*UserContact, error)
	}

	customUserContactModel struct {
		*defaultUserContactModel
	}
)

// NewUserContactModel returns a model for the database table.
func NewUserContactModel(conn sqlx.SqlConn) UserContactModel {
	return &customUserContactModel{
		defaultUserContactModel: newUserContactModel(conn),
	}
}

func (m *customUserContactModel) withSession(session sqlx.Session) UserContactModel {
	return NewUserContactModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserContactModel) Find(ctx context.Context, ownerId int64, req *types.QueryListReq) ([]*UserContact, error) {
	orderField := fmt.Sprintf("`%s`", req.OrderField)
	if !util.Contains(userContactFieldNames, orderField) {
		return nil, &exception.UnknownAscFieldError
	}
	asc := "ASC"
	if !req.Asc {
		asc = "DESC"
	}
	offset := (req.Page - 1) * (req.PageSize)
	var resp []*UserContact
	query := fmt.Sprintf("select %s from %s where OwnerId = ? order by %s %s limit ?,?", userContactRows, m.table, req.OrderField, asc)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, ownerId, offset, req.PageSize)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customUserContactModel) Count(ctx context.Context, ownerId int64) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) as Num from %s where OwnerId = ?", m.table)
	err := m.conn.QueryRowPartialCtx(ctx, &count, query, ownerId)
	switch err {
	case nil:
		return count, nil
	default:
		return -1, err
	}
}

func (m *customUserContactModel) FindDstContact(ctx context.Context, ownerId int64, dstId int64) (*UserContact, error) {
	resp := &UserContact{}
	query := fmt.Sprintf("select %s from %s where OwnerId = ? and DstId = ?", userContactRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, ownerId, dstId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
