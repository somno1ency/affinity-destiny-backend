package group

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		withSession(session sqlx.Session) GroupModel

		CountByOwnerId(ctx context.Context, ownerId int64) (int64, error)
		FindInIds(ctx context.Context, ids []int64) ([]*Group, error)
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

func (m *customGroupModel) CountByOwnerId(ctx context.Context, ownerId int64) (int64, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `owner_id` = ?", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, ownerId)
	switch err {
	case nil:
		return count, nil
	default:
		return 0, err
	}
}

func (m *customGroupModel) FindInIds(ctx context.Context, ids []int64) ([]*Group, error) {
	query := fmt.Sprintf("select %s.* from %s where `id` in (?)", m.table, m.table)
	var resp []*Group
	err := m.conn.QueryRowCtx(ctx, &resp, query, ids)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
