package group_contact

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupContactModel = (*customGroupContactModel)(nil)

type (
	// GroupContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupContactModel.
	GroupContactModel interface {
		groupContactModel
		withSession(session sqlx.Session) GroupContactModel

		FindUserIdsByGroupId(ctx context.Context, id int64) ([]int64, error)
		FindGroupIdsByUserId(ctx context.Context, userId int64) ([]int64, error)
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

func (m *customGroupContactModel) FindUserIdsByGroupId(ctx context.Context, id int64) ([]int64, error) {
	query := fmt.Sprintf("select `user_id` from %s where `group_id` = ?", m.table)
	var resp []int64
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return resp, nil
	default:
		return []int64{}, err
	}
}

func (m *customGroupContactModel) FindGroupIdsByUserId(ctx context.Context, userId int64) ([]int64, error) {
	query := fmt.Sprintf("select `group_id` from %s where `user_id` = ?", m.table)
	var resp []int64
	err := m.conn.QueryRowsCtx(ctx, &resp, query, userId)
	switch err {
	case nil:
		return resp, nil
	default:
		return []int64{}, err
	}
}
