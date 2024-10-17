// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package category

import (
	"context"
	"fmt"

	"ad.com/pkg/exception"
	"ad.com/pkg/shared/types"
	"ad.com/pkg/util"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CategoryModel = (*customCategoryModel)(nil)

type (
	// CategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCategoryModel.
	CategoryModel interface {
		categoryModel
		withSession(session sqlx.Session) CategoryModel

		Find(ctx context.Context, ownerId int64, req *types.QueryListReq) ([]*Category, error)
		Count(ctx context.Context, ownerId int64) (int64, error)
	}

	customCategoryModel struct {
		*defaultCategoryModel
	}
)

// NewCategoryModel returns a model for the database table.
func NewCategoryModel(conn sqlx.SqlConn) CategoryModel {
	return &customCategoryModel{
		defaultCategoryModel: newCategoryModel(conn),
	}
}

func (m *customCategoryModel) withSession(session sqlx.Session) CategoryModel {
	return NewCategoryModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customCategoryModel) Find(ctx context.Context, ownerId int64, req *types.QueryListReq) ([]*Category, error) {
	orderField := fmt.Sprintf("`%s`", req.OrderField)
	if !util.Contains(categoryFieldNames, orderField) {
		return nil, &exception.UnknownAscFieldError
	}
	asc := "ASC"
	if !req.Asc {
		asc = "DESC"
	}
	offset := (req.Page - 1) * (req.PageSize)
	var resp []*Category

	query := fmt.Sprintf("select %s from %s where owner_id = ? order by %s %s limit ?,?", categoryRows, m.table, req.OrderField, asc)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, ownerId, offset, req.PageSize)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *customCategoryModel) Count(ctx context.Context, ownerId int64) (int64, error) {
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
