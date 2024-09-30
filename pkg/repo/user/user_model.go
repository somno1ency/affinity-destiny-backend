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

		FindInIds(ctx context.Context, ids []int64) ([]*User, error)
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

func (m *customUserModel) FindInIds(ctx context.Context, ids []int64) ([]*User, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (?)", userRows, m.table)
	var resp []*User
	err := m.conn.QueryRowsCtx(ctx, &resp, query, ids)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}
