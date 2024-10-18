// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package resource

import (
	"context"
	"fmt"

	"ad.com/pkg/exception"
	"ad.com/pkg/shared/types"
	"ad.com/pkg/util"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ResourceModel = (*customResourceModel)(nil)

type (
	// ResourceModel is an interface to be customized, add more methods here,
	// and implement the added methods in customResourceModel.
	ResourceModel interface {
		resourceModel
		withSession(session sqlx.Session) ResourceModel

		Find(ctx context.Context, resourceType int64, req *types.QueryListReq) ([]*Resource, error)
		Count(ctx context.Context, resourceType int64) (int64, error)
	}

	customResourceModel struct {
		*defaultResourceModel
	}
)

// NewResourceModel returns a model for the database table.
func NewResourceModel(conn sqlx.SqlConn) ResourceModel {
	return &customResourceModel{
		defaultResourceModel: newResourceModel(conn),
	}
}

func (m *customResourceModel) withSession(session sqlx.Session) ResourceModel {
	return NewResourceModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customResourceModel) Find(ctx context.Context, resourceType int64, req *types.QueryListReq) ([]*Resource, error) {
	orderField := fmt.Sprintf("`%s`", req.OrderField)
	if !util.Contains(resourceFieldNames, orderField) {
		return nil, &exception.UnknownAscFieldError
	}
	asc := "ASC"
	if !req.Asc {
		asc = "DESC"
	}
	offset := (req.Page - 1) * (req.PageSize)
	var query string
	var resp []*Resource
	var err error
	if resourceType == 0 {
		query = fmt.Sprintf("select %s from %s order by %s %s limit ?,?", resourceRows, m.table, req.OrderField, asc)
		err = m.conn.QueryRowsCtx(ctx, &resp, query, offset, req.PageSize)
	} else {
		query = fmt.Sprintf("select %s from %s where Type = ? order by %s %s limit ?,?", resourceRows, m.table, req.OrderField, asc)
		err = m.conn.QueryRowsCtx(ctx, &resp, query, resourceType, offset, req.PageSize)
	}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customResourceModel) Count(ctx context.Context, resourceType int64) (int64, error) {
	var query string
	var count int64
	var err error
	if resourceType == 0 {
		query = fmt.Sprintf("select count(*) as Num from %s", m.table)
		err = m.conn.QueryRowPartialCtx(ctx, &count, query)
	} else {
		query = fmt.Sprintf("select count(*) as Num from %s where Type = ?", m.table)
		err = m.conn.QueryRowPartialCtx(ctx, &count, query, resourceType)
	}
	switch err {
	case nil:
		return count, nil
	default:
		return -1, err
	}
}
