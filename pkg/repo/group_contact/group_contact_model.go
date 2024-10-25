// Copyright 2024 Mackay Zhou <mackay.chow@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package group_contact

import (
	"context"
	"fmt"

	"ad.com/pkg/shared"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupContactModel = (*customGroupContactModel)(nil)

type (
	// GroupContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupContactModel.
	GroupContactModel interface {
		groupContactModel
		withSession(session sqlx.Session) GroupContactModel

		Find(ctx context.Context, groupId int64) ([]*GroupContact, error)
		FindByUserId(ctx context.Context, groupId int64, userId int64) (*GroupContact, error)
		BatchInsert(ctx context.Context, data []*GroupContact) error
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

func (m *customGroupContactModel) FindByUserId(ctx context.Context, groupId int64, userId int64) (*GroupContact, error) {
	var resp GroupContact
	query := fmt.Sprintf("select %s from %s where GroupId = ? and UserId = ? limit 1", groupContactRows, m.table)
	err := m.conn.QueryRowCtx(ctx, &resp, query, groupId, userId)
	switch err {
	case nil:
		return &resp, nil
	default:
		return nil, err
	}
}

func (m *customGroupContactModel) BatchInsert(ctx context.Context, data []*GroupContact) error {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, "GroupId, UserId, CreatedAt")
	blk, err := sqlx.NewBulkInserter(m.conn, query)
	if err != nil {
		return err
	}
	for _, v := range data {
		blk.Insert(v.GroupId.Int64, v.UserId.Int64, v.CreatedAt.Time.Local().Format(shared.TimeFormatTemplate))
	}
	blk.Flush()

	return nil
}

func (m *customGroupContactModel) Find(ctx context.Context, groupId int64) ([]*GroupContact, error) {
	var resp []*GroupContact
	query := fmt.Sprintf("select %s from %s where GroupId = ?", groupContactRows, m.table)
	err := m.conn.QueryRowsCtx(ctx, &resp, query, groupId)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
