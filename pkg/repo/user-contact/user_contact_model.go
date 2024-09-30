package user_contact

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserContactModel = (*customUserContactModel)(nil)

type (
	// UserContactModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserContactModel.
	UserContactModel interface {
		userContactModel
		withSession(session sqlx.Session) UserContactModel

		IsContact(ctx context.Context, ownerId int64, dstId int64) (bool, error)
		FindDstIdsByOwnerId(ctx context.Context, ownerId int64) ([]int64, error)
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

func (m *customUserContactModel) IsContact(ctx context.Context, ownerId int64, dstId int64) (bool, error) {
	query := fmt.Sprintf("select count(*) as count from %s where `owner_id` = ? and `dst_id` = ?", m.table)
	var count int64
	err := m.conn.QueryRowCtx(ctx, &count, query, ownerId, dstId)
	switch err {
	case nil:
		return count > 0, nil
	default:
		return false, err
	}
}

func (m *customUserContactModel) FindDstIdsByOwnerId(ctx context.Context, ownerId int64) ([]int64, error) {
	query := fmt.Sprintf("select `dst_id` from %s where `owner_id` = ?", m.table)
	var dstIds []int64
	err := m.conn.QueryRowsCtx(ctx, &dstIds, query, ownerId)
	switch err {
	case nil:
		return dstIds, nil
	default:
		return []int64{}, err
	}
}
