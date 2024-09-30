// Code generated by goctl. DO NOT EDIT.

package group_contact

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupContactFieldNames          = builder.RawFieldNames(&GroupContact{})
	groupContactRows                = strings.Join(groupContactFieldNames, ",")
	groupContactRowsExpectAutoSet   = strings.Join(stringx.Remove(groupContactFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupContactRowsWithPlaceHolder = strings.Join(stringx.Remove(groupContactFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	groupContactModel interface {
		Insert(ctx context.Context, data *GroupContact) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*GroupContact, error)
		Update(ctx context.Context, data *GroupContact) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGroupContactModel struct {
		conn  sqlx.SqlConn
		table string
	}

	GroupContact struct {
		Id        int64         `db:"id"`
		GroupId   sql.NullInt64 `db:"group_id"`   // 群ID
		UserId    sql.NullInt64 `db:"user_id"`    // 用户ID
		CreatedAt sql.NullTime  `db:"created_at"` // 创建时间
	}
)

func newGroupContactModel(conn sqlx.SqlConn) *defaultGroupContactModel {
	return &defaultGroupContactModel{
		conn:  conn,
		table: "`group_contact`",
	}
}

func (m *defaultGroupContactModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultGroupContactModel) FindOne(ctx context.Context, id int64) (*GroupContact, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupContactRows, m.table)
	var resp GroupContact
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupContactModel) Insert(ctx context.Context, data *GroupContact) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, groupContactRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.GroupId, data.UserId)
	return ret, err
}

func (m *defaultGroupContactModel) Update(ctx context.Context, data *GroupContact) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupContactRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.GroupId, data.UserId, data.Id)
	return err
}

func (m *defaultGroupContactModel) tableName() string {
	return m.table
}
