// Code generated by goctl. DO NOT EDIT.

package category

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
	categoryFieldNames          = builder.RawFieldNames(&Category{})
	categoryRows                = strings.Join(categoryFieldNames, ",")
	categoryRowsExpectAutoSet   = strings.Join(stringx.Remove(categoryFieldNames, "`id`"), ",")
	categoryRowsWithPlaceHolder = strings.Join(stringx.Remove(categoryFieldNames, "`id`"), "=?,") + "=?"
)

type (
	categoryModel interface {
		Insert(ctx context.Context, data *Category) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Category, error)
		Update(ctx context.Context, data *Category) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCategoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Category struct {
		Id        int64        `db:"id"`
		OwnerId   int64        `db:"owner_id"`   // 创建者ID
		NameZh    string       `db:"name_zh"`    // 分类中文名称
		NameEn    string       `db:"name_en"`    // 分类英文名称
		Star      int64        `db:"star"`       // 分类星级
		CreatedAt sql.NullTime `db:"created_at"` // 创建时间
		UpdatedAt sql.NullTime `db:"updated_at"` // 更新时间
	}
)

func newCategoryModel(conn sqlx.SqlConn) *defaultCategoryModel {
	return &defaultCategoryModel{
		conn:  conn,
		table: "`category`",
	}
}

func (m *defaultCategoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultCategoryModel) FindOne(ctx context.Context, id int64) (*Category, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", categoryRows, m.table)
	var resp Category
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

func (m *defaultCategoryModel) Insert(ctx context.Context, data *Category) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, categoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OwnerId, data.NameZh, data.NameEn, data.Star, data.CreatedAt, data.UpdatedAt)
	return ret, err
}

func (m *defaultCategoryModel) Update(ctx context.Context, data *Category) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, categoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OwnerId, data.NameZh, data.NameEn, data.Star, data.CreatedAt, data.UpdatedAt, data.Id)
	return err
}

func (m *defaultCategoryModel) tableName() string {
	return m.table
}
