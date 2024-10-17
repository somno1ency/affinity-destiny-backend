// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group

import (
	"context"
	"fmt"

	"ad.com/pkg/exception"
	"ad.com/pkg/shared/types"
	"ad.com/pkg/util"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		withSession(session sqlx.Session) GroupModel

		Find(ctx context.Context, ownerId int64, req *types.QueryListReq) ([]*Group, error)
		Count(ctx context.Context, ownerId int64) (int64, error)
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn),
	}
}

func (m *customGroupModel) withSession(session sqlx.Session) GroupModel {
	return NewGroupModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customGroupModel) Find(ctx context.Context, ownerId int64, req *types.QueryListReq) ([]*Group, error) {
	orderField := fmt.Sprintf("`%s`", req.OrderField)
	if !util.Contains(groupFieldNames, orderField) {
		return nil, &exception.UnknownAscFieldError
	}
	asc := "ASC"
	if !req.Asc {
		asc = "DESC"
	}
	offset := (req.Page - 1) * (req.PageSize)
	var resp []*Group

	query := fmt.Sprintf("select %s from %s where owner_id = ? order by %s %s limit ?,?", groupRows, m.table, req.OrderField, asc)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, ownerId, offset, req.PageSize)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customGroupModel) Count(ctx context.Context, ownerId int64) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) as num from %s where owner_id = ?", m.table)
	err := m.conn.QueryRowPartialCtx(ctx, &count, query, ownerId)

	switch err {
	case nil:
		return count, nil
	default:
		return -1, err
	}
}
